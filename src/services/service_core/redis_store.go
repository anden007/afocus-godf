package service_core

import (
	"time"

	"github.com/anden007/afocus-godf/src/interfaces/interface_core"
)

type RedisStore struct {
	cache interface_core.ICache
}

func NewRedisStore(cache interface_core.ICache) (store *RedisStore) {
	store = new(RedisStore)
	store.cache = cache
	return store
}

func (m *RedisStore) Get(id string, clear bool) string {
	result, err := m.cache.Get(id)
	if err == nil {
		if clear {
			_, _ = m.cache.Del(id)
		}
	}
	return result
}

func (m *RedisStore) Set(id string, value string) {
	_, _ = m.cache.SetEx(id, value, time.Minute)
}

func (m *RedisStore) Verify(id, answer string, clear bool) bool {
	result := false
	value, err := m.cache.Get(id)
	if err == nil {
		result = answer == value
		if clear {
			_, _ = m.cache.Del(id)
		}
	}
	return result
}
