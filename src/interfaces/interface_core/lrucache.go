package interface_core

import "time"

type ILruCache interface {
	Add(key, value interface{}) (success bool)
	AddEx(key, value interface{}, expire time.Duration) (success bool)
	Get(key interface{}) (result interface{}, success bool)
	Contains(key interface{}) (success bool)
	Peek(key interface{}) (result interface{}, success bool)
	ContainsOrAdd(key, value interface{}) (success, evict bool)
	Remove(key interface{})
	RemoveOldest()
	Keys() (result []interface{})
	Len() (result int)
	Purge()
}
