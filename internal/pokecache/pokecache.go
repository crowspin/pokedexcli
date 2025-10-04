package pokecache

import (
	"sync"
	"time"
)

const (
	MAX_CACHE_SECONDS time.Duration = 10 * time.Second
)

type Cache struct {
	internal map[string]cacheEntry
	mu       *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		internal: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.internal[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.internal[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.internal {
			if time.Now().Add(-MAX_CACHE_SECONDS).After(v.createdAt) {
				delete(c.internal, k)
			}
		}
		c.mu.Unlock()
	}
}
