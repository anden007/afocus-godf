// This package provides a simple LRU cache. It is based on the
// LRU implementation in groupcache:
// https://github.com/golang/groupcache/tree/master/lru
package service_core

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/anden007/afocus-godf/src/lib"
	"github.com/anden007/afocus-godf/src/services/service_core/simplelru"

	"github.com/spf13/cast"
)

// Cache is a thread-safe fixed size LRU cache.
type LruCache struct {
	lru  *simplelru.LRU
	lock sync.RWMutex
}

func NewLruCache() *LruCache {
	var instance *LruCache
	loadTime := time.Now()
	lruCacheSize := cast.ToInt(os.Getenv("lru-cache-size"))
	lruCacheExpire := cast.ToInt(os.Getenv("lru-cache-expire"))
	instance, _ = NewWithExpire(lruCacheSize, time.Second*time.Duration(lruCacheExpire))
	if lib.IS_DEV_MODE {
		fmt.Println("> Service: LruCache loaded.", time.Since(loadTime))
	}
	return instance
}

// New creates an LRU of the given size
func New(size int) (*LruCache, error) {
	return NewWithEvict(size, nil)
}

// NewWithEvict constructs a fixed size cache with the given eviction
// callback.
func NewWithEvict(size int, onEvicted func(key interface{}, value interface{})) (*LruCache, error) {
	lru, err := simplelru.NewLRU(size, simplelru.EvictCallback(onEvicted))
	if err != nil {
		return nil, err
	}
	c := &LruCache{
		lru: lru,
	}
	return c, nil
}

// NewWithExpire constructs a fixed size cache with expire feature
func NewWithExpire(size int, expire time.Duration) (*LruCache, error) {
	lru, err := simplelru.NewLRUWithExpire(size, expire, nil)
	if err != nil {
		return nil, err
	}
	c := &LruCache{
		lru: lru,
	}
	return c, nil
}

// Purge is used to completely clear the cache
func (c *LruCache) Purge() {
	c.lock.Lock()
	c.lru.Purge()
	c.lock.Unlock()
}

// Add adds a value to the cache.  Returns true if an eviction occurred.
func (c *LruCache) Add(key, value interface{}) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.lru.Add(key, value)
}

// AddEx adds a value to the cache.  Returns true if an eviction occurred.
func (c *LruCache) AddEx(key, value interface{}, expire time.Duration) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.lru.AddEx(key, value, expire)
}

// Get looks up a key's value from the cache.
func (c *LruCache) Get(key interface{}) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.lru.Get(key)
}

// Check if a key is in the cache, without updating the recent-ness
// or deleting it for being stale.
func (c *LruCache) Contains(key interface{}) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.lru.Contains(key)
}

// Returns the key value (or undefined if not found) without updating
// the "recently used"-ness of the key.
func (c *LruCache) Peek(key interface{}) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.lru.Peek(key)
}

// ContainsOrAdd checks if a key is in the cache  without updating the
// recent-ness or deleting it for being stale,  and if not, adds the value.
// Returns whether found and whether an eviction occurred.
func (c *LruCache) ContainsOrAdd(key, value interface{}) (ok, evict bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.lru.Contains(key) {
		return true, false
	} else {
		evict := c.lru.Add(key, value)
		return false, evict
	}
}

// Remove removes the provided key from the cache.
func (c *LruCache) Remove(key interface{}) {
	c.lock.Lock()
	c.lru.Remove(key)
	c.lock.Unlock()
}

// RemoveOldest removes the oldest item from the cache.
func (c *LruCache) RemoveOldest() {
	c.lock.Lock()
	c.lru.RemoveOldest()
	c.lock.Unlock()
}

// Keys returns a slice of the keys in the cache, from oldest to newest.
func (c *LruCache) Keys() []interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.lru.Keys()
}

// Len returns the number of items in the cache.
func (c *LruCache) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.lru.Len()
}
