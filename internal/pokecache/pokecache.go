package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	internal map[string]cacheEntry
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	return Cache{}
}

func (c Cache) Add(key string, val []byte) {

}

func (c Cache) Get(key string) ([]byte, bool) {
	return nil, false
}

func (c Cache) reapLoop() {

}
