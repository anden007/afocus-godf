package interface_core

import (
	"time"

	"github.com/kataras/iris/v12"
)

type ISessionCache interface {
	Get(ctx iris.Context, key string) (result string, err error)
	Set(ctx iris.Context, key string, value string, ttl time.Duration) (result string, err error)
	Del(ctx iris.Context, key string) (result bool, err error)
	Exists(ctx iris.Context, key string) (result bool, err error)
	Incr(ctx iris.Context, key string) (result int64, err error)
	GetCacheKey(ctx iris.Context, key string) string
	Md5(str string) string
}
