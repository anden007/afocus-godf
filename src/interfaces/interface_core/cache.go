package interface_core

import (
	"time"

	"github.com/go-redis/redis"
)

type ICache interface {
	GetBackend() (result *redis.Client, err error)
	Get(key string) (result string, err error)
	Set(key string, value string) (result string, err error)
	SetEx(key string, value string, ttl time.Duration) (result string, err error)
	Del(key string) (result bool, err error)
	Exists(key string) (result bool, err error)
	Incr(key string) (result int64, err error)
	Decr(key string) (result int64, err error)
	CreateKey(key string) string
}
