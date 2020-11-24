package service_core

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/model/model_draw"

	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	jsontime "github.com/liamylian/jsontime/v2/v2"
	"github.com/spf13/cast"
)

type DrawCenter struct {
	Enable bool
	redis  *redis.Client
	json   jsoniter.API
	lock   *sync.RWMutex
}

type CacheKeyType int32

const (
	ACTION_KEY                  CacheKeyType = 1 // 活动Key
	ACTION_WINNER_LIST_KEY      CacheKeyType = 2 // 活动中奖者列表Key
	ACTION_PLAYER_KEY           CacheKeyType = 3 // 活动用户Key
	ACTION_PLAYER_GIFT_LIST_KEY CacheKeyType = 4 // 活动用户获得的奖品列表Key
	ACTION_GIFT_LIST_KEY        CacheKeyType = 5 // 活动奖品列表Key
	ACTION_GIFT_KEY             CacheKeyType = 6 // 活动奖品Key
)

func NewDrawCenter() *DrawCenter {
	instance := new(DrawCenter)
	loadTime := time.Now()
	enable := os.Getenv("draw-enable")
	if strings.ToLower(enable) == "true" {
		instance.Enable = true
		instance.json = jsontime.ConfigWithCustomTimeFormat
		instance.lock = new(sync.RWMutex)
		db := cast.ToInt(os.Getenv("draw-db"))
		poolSize := cast.ToInt(os.Getenv("redis-pool-size"))
		instance.redis = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("redis-server"),
			Password: os.Getenv("redis-password"),
			DB:       db,
			PoolSize: poolSize,
		})
		if lib.IS_DEV_MODE {
			fmt.Println("> Service: DrawCenter loaded.", time.Since(loadTime))
		}
	} else {
		fmt.Printf("> Service: DrawCenter is disabled. if you need enable it,please set 'draw-enable' to 'true' in .env file.\n")
	}
	return instance
}

// 执行抽奖
func (m *DrawCenter) Draw(actionId string, playerId string) (result model_draw.Result, err error) {
	if !m.Enable {
		panic("DrawCenter is disabled. if you need enable it,please set 'draw-enable' to 'true' in .env file.\n")
	}
	if actionId != "" && playerId != "" {
		playerKey := m.getCacheKey(ACTION_PLAYER_KEY, actionId, playerId)
		if player, err := m.GetPlayer(actionId, playerId); err == nil {
			// 锁不能放到GetPlayer之上，否则会出现死锁现相，因为GetPlayer方法内也有锁定操作
			m.lock.Lock()
			defer m.lock.Unlock()
			if action, err := m.GetAction(actionId); err == nil {
				if player.RemainPlayCount <= 0 {
					result = model_draw.Result{
						ActionID: actionId,
						PlayerID: playerId,
						GiftType: 0,
						GiftID:   "",
						GiftName: "",
						Voucher:  nil,
						Player:   &player,
						Success:  false,
						Message:  "已达到最大可抽奖次数",
					}
					err = errors.New("已达到最大可抽奖次数")
					return result, err
				}
				if player.RemainDayPlayCount <= 0 {
					result = model_draw.Result{
						ActionID: actionId,
						PlayerID: playerId,
						GiftType: 0,
						GiftID:   "",
						GiftName: "",
						Voucher:  nil,
						Player:   &player,
						Success:  false,
						Message:  "已达到今日最大可抽奖次数",
					}
					err = errors.New("已达到今日最大可抽奖次数")
					return result, err
				}
				if player.RemainWinCount <= 0 {
					result = model_draw.Result{
						ActionID: actionId,
						PlayerID: playerId,
						GiftType: 0,
						GiftID:   "",
						GiftName: "",
						Voucher:  nil,
						Player:   &player,
						Success:  false,
						Message:  "已达到最大可中奖次数",
					}
					err = errors.New("已达到最大可中奖次数")
					return result, err
				}
				if player.RemainDayWinCount <= 0 {
					result = model_draw.Result{
						ActionID: actionId,
						PlayerID: playerId,
						GiftType: 0,
						GiftID:   "",
						GiftName: "",
						Voucher:  nil,
						Player:   &player,
						Success:  false,
						Message:  "已达到今日最大可中奖次数",
					}
					err = errors.New("已达到今日最大可中奖次数")
					return result, err
				}
				if gifts, err := m.GetActionGifts(actionId); err == nil {
					if action.Enable == 1 && action.BeginTime.Before(time.Now()) && action.EndTime.After(time.Now()) {
						// 活动开启且在开放时间段内
						var toDrawGifts []model_draw.Gift
						var fullChanceGifts []model_draw.Gift
						for _, gift := range gifts {
							if gift.Enable == 1 && gift.BeginTime.Before(time.Now()) && gift.EndTime.After(time.Now()) && gift.RemainCount > 0 {
								if gift.Odds == 1000000 {
									// 收集100%中奖的奖品
									fullChanceGifts = append(fullChanceGifts, gift)
								}
								//收集可用于抽奖的奖品
								toDrawGifts = append(toDrawGifts, gift)
							}
						}
						if len(toDrawGifts) > 0 {
							// 更新用户抽奖相关次数
							player.PlayCount++
							player.RemainPlayCount--
							player.DayPlayCount++
							player.RemainDayPlayCount--
							player.LastPlayTime = time.Now()
							pip := m.redis.Pipeline()
							pip.HIncrBy(playerKey, "PlayCount", 1)
							pip.HIncrBy(playerKey, "DayPlayCount", 1)
							pip.HSet(playerKey, "LastPlayTime", time.Now().Format("2006-01-02 15:04:05"))
							if _, err = pip.Exec(); err == nil {
								var targetGift model_draw.Gift
								targetGiftIndex := 0
								if len(toDrawGifts) > 1 {
									rand.Seed(time.Now().UnixNano())
									targetGiftIndex = int(math.Trunc(rand.Float64() * float64(len(toDrawGifts))))
								}
								rand.Seed(time.Now().UnixNano())
								userRand := rand.Intn(1000000)
								if toDrawGifts[targetGiftIndex].Odds >= userRand {
									// 按随机数抽取奖品
									// 设置奖品消耗及用户抽奖信息
									// 更新用户中奖相关次数
									targetGift = toDrawGifts[targetGiftIndex]
								} else if len(fullChanceGifts) > 0 {
									// 处理未中奖但有100%概率的奖品的情况
									rand.Seed(time.Now().UnixNano())
									targetGiftIndex := 0
									if len(fullChanceGifts) > 1 {
										targetGiftIndex = rand.Intn(len(fullChanceGifts) - 1)
									}
									targetGift = fullChanceGifts[targetGiftIndex]
								} else {
									result = model_draw.Result{
										ActionID: actionId,
										PlayerID: playerId,
										GiftType: 0,
										GiftID:   "",
										GiftName: "",
										Voucher:  nil,
										Player:   &player,
										Success:  false,
										Message:  "很遗憾，未中奖",
									}
									return result, nil
								}
								giftKey := m.getCacheKey(ACTION_GIFT_KEY, actionId, targetGift.Id)
								userGiftKey := m.getCacheKey(ACTION_PLAYER_GIFT_LIST_KEY, actionId, playerId)
								winnerListKey := m.getCacheKey(ACTION_WINNER_LIST_KEY, actionId)
								player.WinCount++
								player.RemainWinCount--
								player.DayWinCount++
								player.RemainDayWinCount--
								targetGift.RemainCount--
								pip := m.redis.Pipeline()
								pip.RPush(userGiftKey, targetGift.Id)
								pip.RPush(winnerListKey, playerId)
								pip.HIncrBy(playerKey, "WinCount", 1)
								pip.HIncrBy(playerKey, "DayWinCount", 1)
								pip.HIncrBy(giftKey, "RemainCount", -1)
								if _, err = pip.Exec(); err == nil {
									result = model_draw.Result{
										ActionID: actionId,
										PlayerID: playerId,
										GiftType: targetGift.GiftType,
										GiftID:   targetGift.Id,
										GiftName: targetGift.GiftName,
										Voucher:  nil,
										Player:   &player,
										Success:  true,
										Message:  "",
									}
								}
							}
						} else {
							result = model_draw.Result{
								ActionID: actionId,
								PlayerID: playerId,
								GiftType: 0,
								GiftID:   "",
								GiftName: "",
								Voucher:  nil,
								Player:   &player,
								Success:  false,
								Message:  "很遗憾，未中奖",
							}
							err = errors.New("无候选奖品")
							return result, err
						}
					} else if action.Enable == 0 {
						return model_draw.Result{}, errors.New("活动已关闭")
					} else if action.BeginTime.After(time.Now()) {
						return model_draw.Result{}, errors.New("活动尚未开始")
					} else if action.EndTime.Before(time.Now()) {
						return model_draw.Result{}, errors.New("活动已结束")
					}
				}
			}
		}
	} else if actionId == "" {
		err = errors.New("Error: ActionId is needed ")
	} else if playerId == "" {
		err = errors.New("Error: PlayerId is needed ")
	}
	return
}

// 用户确认获奖
func (m *DrawCenter) Verify(actionId string, playerId string, giftId string, verify bool) (result model_draw.VerifyResult, err error) {
	if !m.Enable {
		panic("DrawCenter is disabled. if you need enable it,please set 'draw-enable' to 'true' in .env file.\n")
	}
	return
}

// 获得活动信息
func (m *DrawCenter) GetAction(actionId string) (result model_draw.Action, err error) {
	if !m.Enable {
		panic("DrawCenter is disabled. if you need enable it,please set 'draw-enable' to 'true' in .env file.\n")
	}
	if actionId != "" {
		key := m.getCacheKey(ACTION_KEY, actionId)
		redisResult := m.redis.HGetAll(key)
		if err = redisResult.Err(); err == nil {
			resultMap := redisResult.Val()
			enable := cast.ToInt(resultMap["Enable"])
			beginTime, _ := time.ParseInLocation("2006-01-02 15:04:05", resultMap["BeginTime"], time.Local)
			endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", resultMap["EndTime"], time.Local)
			maxPlayCount := cast.ToInt(resultMap["MaxPlayCount"])
			maxWinCount := cast.ToInt(resultMap["MaxWinCount"])
			dayMaxPlayCount := cast.ToInt(resultMap["DayMaxPlayCount"])
			dayMaxWinCount := cast.ToInt(resultMap["DayMaxWinCount"])
			confirmSeconds := cast.ToInt(resultMap["ConfirmSeconds"])
			result = model_draw.Action{
				Id:              resultMap["Id"],
				Enable:          enable,
				Name:            resultMap["Name"],
				BeginTime:       beginTime,
				EndTime:         endTime,
				MaxPlayCount:    maxPlayCount,
				MaxWinCount:     maxWinCount,
				DayMaxPlayCount: dayMaxPlayCount,
				DayMaxWinCount:  dayMaxWinCount,
				ConfirmSeconds:  confirmSeconds,
			}
		}
	} else {
		err = errors.New("Error: ActionId is needed ")
	}
	return
}

// 获得活动奖品列表
func (m *DrawCenter) GetActionGifts(actionId string) (result []model_draw.Gift, err error) {
	if !m.Enable {
		panic("DrawCenter is disabled. if you need enable it,please set 'draw-enable' to 'true' in .env file.\n")
	}
	if actionId != "" {
		key := m.getCacheKey(ACTION_GIFT_LIST_KEY, actionId)
		cliResult := m.redis.SMembers(key)
		err = cliResult.Err()
		if err == nil {
			giftIds := cliResult.Val()
			for _, giftId := range giftIds {
				if gift, err := m.GetGift(actionId, giftId); err == nil {
					result = append(result, gift)
				}
			}
		}
	} else {
		err = errors.New("Error: ActionId is needed ")
	}
	return
}

// 获得活动奖品信息（单个）
func (m *DrawCenter) GetGift(actionId string, giftId string) (result model_draw.Gift, err error) {
	if !m.Enable {
		panic("DrawCenter is disabled. if you need enable it,please set 'draw-enable' to 'true' in .env file.\n")
	}
	if actionId != "" && giftId != "" {
		key := m.getCacheKey(ACTION_GIFT_KEY, actionId, giftId)
		redisResult := m.redis.HGetAll(key)
		if err = redisResult.Err(); err == nil {
			resultMap := redisResult.Val()
			enable := cast.ToInt(resultMap["Enable"])
			beginTime, _ := time.ParseInLocation("2006-01-02 15:04:05", resultMap["BeginTime"], time.Local)
			endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", resultMap["EndTime"], time.Local)
			giftType := cast.ToInt(resultMap["GiftType"])
			odds := cast.ToInt(resultMap["Odds"])
			count := cast.ToInt(resultMap["Count"])
			remainCount := cast.ToInt(resultMap["RemainCount"])
			result = model_draw.Gift{
				Id:          resultMap["Id"],
				ActionID:    resultMap["ActionID"],
				Enable:      enable,
				GiftType:    giftType,
				GiftName:    resultMap["GiftName"],
				BeginTime:   beginTime,
				EndTime:     endTime,
				Odds:        odds,
				Count:       count,
				RemainCount: remainCount,
			}
		}
	} else if actionId == "" {
		err = errors.New("Error: ActionId is needed ")
	} else if giftId == "" {
		err = errors.New("Error: GiftId is needed ")
	}
	return
}

// 获得活动用户信息
func (m *DrawCenter) GetPlayer(actionId string, playerId string) (result model_draw.Player, err error) {
	if !m.Enable {
		panic("DrawCenter is disabled. if you need enable it,please set 'draw-enable' to 'true' in .env file.\n")
	}
	playerKey := m.getCacheKey(ACTION_PLAYER_KEY, actionId, playerId)
	if action, err := m.GetAction(actionId); err == nil {
		existsResult := m.redis.Exists(playerKey)
		err = existsResult.Err()
		if err == nil && existsResult.Val() == 1 {
			m.lock.RLock()
			redisResult := m.redis.HGetAll(playerKey)
			m.lock.RUnlock()
			if err = redisResult.Err(); err == nil {
				resultMap := redisResult.Val()
				playCount := cast.ToInt(resultMap["PlayCount"])
				dayPlayCount := cast.ToInt(resultMap["DayPlayCount"])
				winCount := cast.ToInt(resultMap["WinCount"])
				dayWinCount := cast.ToInt(resultMap["DayWinCount"])
				lastPlayTime, _ := time.ParseInLocation("2006-01-02 15:04:05", resultMap["LastPlayTime"], time.Local)
				if lastPlayTime.Before(time.Now()) {
					dayPlayCount = 0
					dayWinCount = 0
					m.lock.Lock()
					pip := m.redis.Pipeline()
					pip.HSet(playerKey, "DayPlayCount", 0)
					pip.HSet(playerKey, "DayWinCount", 0)
					_, _ = pip.Exec()
					m.lock.Unlock()
				}
				remainPlayCount := 0
				if action.MaxPlayCount > 0 {
					remainPlayCount = action.MaxPlayCount - playCount
				} else {
					remainPlayCount = 1
				}

				remainDayPlayCount := 0
				if action.DayMaxPlayCount > 0 {
					remainDayPlayCount = action.DayMaxPlayCount - dayPlayCount
				} else {
					remainDayPlayCount = 1
				}

				remainWinCount := 0
				if action.MaxWinCount > 0 {
					remainWinCount = action.MaxWinCount - winCount
				} else {
					remainWinCount = 1
				}

				remainDayWinCount := 0
				if action.DayMaxWinCount > 0 {
					remainDayWinCount = action.DayMaxWinCount - dayPlayCount
				} else {
					remainDayWinCount = 1
				}

				result = model_draw.Player{
					Id:                 resultMap["Id"],
					ActionID:           resultMap["ActionID"],
					PlayCount:          playCount,
					RemainPlayCount:    remainPlayCount,
					WinCount:           winCount,
					RemainWinCount:     remainWinCount,
					DayPlayCount:       dayPlayCount,
					RemainDayPlayCount: remainDayPlayCount,
					DayWinCount:        dayWinCount,
					RemainDayWinCount:  remainDayWinCount,
					LastPlayTime:       lastPlayTime,
				}
			}
		} else {
			result = model_draw.Player{
				Id:                 playerId,
				ActionID:           actionId,
				PlayCount:          0,
				RemainPlayCount:    action.MaxPlayCount,
				WinCount:           0,
				RemainWinCount:     action.MaxWinCount,
				DayPlayCount:       0,
				RemainDayPlayCount: action.DayMaxPlayCount,
				DayWinCount:        0,
				RemainDayWinCount:  action.DayMaxWinCount,
				LastPlayTime:       time.Now(),
			}
			m.lock.Lock()
			m.redis.HMSet(playerKey, map[string]interface{}{
				"Id":           playerId,
				"ActionID":     actionId,
				"PlayCount":    0,
				"WinCount":     0,
				"DayPlayCount": 0,
				"DayWinCount":  0,
				"LastPlayTime": time.Now().Format("2006-01-02 15:04:05"),
			})
			m.lock.Unlock()
		}
	}
	return
}

// 获得用户已获得奖品列表
func (m *DrawCenter) GetPlayerGifts(actionId string, playerId string) (result []model_draw.Gift, err error) {
	if !m.Enable {
		panic("DrawCenter is disabled. if you need enable it,please set 'draw-enable' to 'true' in .env file.\n")
	}
	key := m.getCacheKey(ACTION_PLAYER_GIFT_LIST_KEY, actionId, playerId)
	listLen := m.redis.LLen(key)
	err = listLen.Err()
	if err == nil {
		cliResult := m.redis.LRange(key, 0, listLen.Val())
		err = cliResult.Err()
		if err == nil {
			giftIds := cliResult.Val()
			for _, giftId := range giftIds {
				if gift, err := m.GetGift(actionId, giftId); err == nil {
					result = append(result, gift)
				}
			}
		}
	}
	return
}

func (m *DrawCenter) GetActionWinners(actionId string) (result []model_draw.Player, err error) {
	if !m.Enable {
		panic("DrawCenter is disabled. if you need enable it,please set 'draw-enable' to 'true' in .env file.\n")
	}
	key := m.getCacheKey(ACTION_WINNER_LIST_KEY, actionId)
	listLen := m.redis.LLen(key)
	err = listLen.Err()
	if err == nil {
		cliResult := m.redis.LRange(key, 0, listLen.Val())
		err = cliResult.Err()
		if err == nil {
			playerIds := cliResult.Val()
			for _, playerId := range playerIds {
				if player, err := m.GetPlayer(actionId, playerId); err == nil {
					result = append(result, player)
				}
			}
		}
	}
	return
}

// (内部)获得缓存Key
func (m *DrawCenter) getCacheKey(keyType CacheKeyType, keys ...string) (result string) {
	switch keyType {
	case ACTION_KEY:
		result = fmt.Sprintf("act_%s", keys[0])
	case ACTION_PLAYER_KEY:
		result = fmt.Sprintf("act-player_%s_%s", keys[0], keys[1])
	case ACTION_PLAYER_GIFT_LIST_KEY:
		result = fmt.Sprintf("act-player-gifts_%s_%s", keys[0], keys[1])
	case ACTION_GIFT_LIST_KEY:
		result = fmt.Sprintf("act-gifts_%s", keys[0])
	case ACTION_GIFT_KEY:
		result = fmt.Sprintf("act-gift_%s_%s", keys[0], keys[1])
	case ACTION_WINNER_LIST_KEY:
		result = fmt.Sprintf("act-winners_%s", keys[0])

	}
	return
}
