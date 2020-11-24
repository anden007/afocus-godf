package service_core

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/anden007/afocus-godf/src/lib"

	"github.com/go-redis/redis"
	"github.com/spf13/cast"
)

type CacheConfig struct {
	AppName  string
	Server   string
	Password string
	DataBase int
	PoolSize int
}

type Cache struct {
	appName string
	redis   *redis.Client
}

func NewCache() *Cache {
	instance := new(Cache)
	loadTime := time.Now()
	db := cast.ToInt(os.Getenv("redis-db"))
	poolSize := cast.ToInt(os.Getenv("redis-pool-size"))
	instance.appName = os.Getenv("app-name")
	instance.redis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis-server"),
		Password: os.Getenv("redis-password"),
		DB:       db,
		PoolSize: poolSize,
	})
	if lib.IS_DEV_MODE {
		fmt.Println("> Service: Cache loaded, ", time.Since(loadTime))
	}
	return instance
}

func NewCacheByConfig(config CacheConfig) *Cache {
	instance := new(Cache)
	instance.appName = config.AppName
	instance.redis = redis.NewClient(&redis.Options{
		Addr:     config.Server,
		Password: config.Password,
		DB:       config.DataBase,
		PoolSize: config.PoolSize,
	})
	return instance
}

func (m *Cache) GetBackend() (result *redis.Client, err error) {
	return m.redis, nil
}

func (m *Cache) Get(key string) (result string, err error) {
	_key := m.CreateKey(key)
	result, err = m.redis.Get(_key).Result()
	return
}

func (m *Cache) Set(key string, value string) (result string, err error) {
	_key := m.CreateKey(key)
	result, err = m.redis.Set(_key, value, 0).Result()
	return
}

func (m *Cache) SetEx(key string, value string, ttl time.Duration) (result string, err error) {
	_key := m.CreateKey(key)
	result, err = m.redis.Set(_key, value, ttl).Result()
	return
}

func (m *Cache) Del(key string) (result bool, err error) {
	_key := m.CreateKey(key)
	cmdResult, err := m.redis.Del(_key).Result()
	return cmdResult > 0, err
}

func (m *Cache) Exists(key string) (result bool, err error) {
	_key := m.CreateKey(key)
	cmdResult, err := m.redis.Exists(_key).Result()
	return cmdResult > 0, err
}

func (m *Cache) Incr(key string) (result int64, err error) {
	_key := m.CreateKey(key)
	cmdResult, err := m.redis.Incr(_key).Result()
	return cmdResult, err
}

func (m *Cache) Decr(key string) (result int64, err error) {
	_key := m.CreateKey(key)
	cmdResult, err := m.redis.Decr(_key).Result()
	return cmdResult, err
}

func (m *Cache) Md5(str string) string {
	h := md5.New()
	_, _ = h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func (m *Cache) CreateKey(key string) string {
	return fmt.Sprintf("%s_%s", m.appName, key)
}
