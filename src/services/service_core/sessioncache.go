package service_core

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/anden007/afocus-godf/src/interfaces/interface_core"

	uuid "github.com/iris-contrib/go.uuid"
	"github.com/kataras/iris/v12"
)

type SessionCache struct {
	sessionId string
	cache     interface_core.ICache
}

func NewSessionCache(cache interface_core.ICache) *SessionCache {
	instance := new(SessionCache)
	instance.cache = cache
	instance.sessionId = instance.Md5(os.Getenv("app-name"))
	return instance
}

func (m *SessionCache) Get(ctx iris.Context, key string) (result string, err error) {
	_key := m.GetCacheKey(ctx, key)
	result, err = m.cache.Get(_key)
	return
}

func (m *SessionCache) Set(ctx iris.Context, key string, value string, ttl time.Duration) (result string, err error) {
	_key := m.GetCacheKey(ctx, key)
	result, err = m.cache.SetEx(_key, value, ttl)
	return
}

func (m *SessionCache) Del(ctx iris.Context, key string) (result bool, err error) {
	_key := m.GetCacheKey(ctx, key)
	result, err = m.cache.Del(_key)
	return
}

func (m *SessionCache) Exists(ctx iris.Context, key string) (result bool, err error) {
	_key := m.GetCacheKey(ctx, key)
	result, err = m.cache.Exists(_key)
	return
}

func (m *SessionCache) Incr(ctx iris.Context, key string) (result int64, err error) {
	_key := m.GetCacheKey(ctx, key)
	result, err = m.cache.Incr(_key)
	return
}

func (m *SessionCache) GetCacheKey(ctx iris.Context, key string) string {
	uid := ctx.GetCookie(m.sessionId)
	if uid == "" {
		newUUID, _ := uuid.NewV4()
		uid = newUUID.String()
		ctx.SetCookieKV(m.sessionId, uid, iris.CookieExpires(time.Hour*24))
	}
	return m.Md5(fmt.Sprintf("%s%s", uid, key))
}

func (m *SessionCache) Md5(str string) string {
	result := ""
	h := md5.New()
	_, err := h.Write([]byte(str))
	if err == nil {
		result = hex.EncodeToString(h.Sum(nil))
	}
	return result
}
