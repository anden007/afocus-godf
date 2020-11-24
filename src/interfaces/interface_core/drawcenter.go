package interface_core

import (
	"github.com/anden007/afocus-godf/src/model/model_draw"
)

type IDrawCenter interface {
	Draw(actionId string, playerId string) (result model_draw.Result, err error)
	Verify(actionId string, playerId string, giftId string, verify bool) (result model_draw.VerifyResult, err error)
	GetAction(actionId string) (result model_draw.Action, err error)
	GetActionGifts(actionId string) (result []model_draw.Gift, err error)
	GetActionWinners(actionId string) (result []model_draw.Player, err error)
	GetGift(actionId string, giftId string) (result model_draw.Gift, err error)
	GetPlayer(actionId string, playerId string) (result model_draw.Player, err error)
	GetPlayerGifts(actionId string, playerId string) (result []model_draw.Gift, err error)
}
