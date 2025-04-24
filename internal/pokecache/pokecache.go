package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]cacheEntry
	mu      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		Entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}
	go newCache.reapLoop(interval)
	return newCache
}

func (ca *Cache) Add(key string, val []byte) {
	ca.mu.Lock()
	defer ca.mu.Unlock()

	ca.Entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (ca *Cache) Get(key string) ([]byte, bool) {
	ca.mu.Lock()
	defer ca.mu.Unlock()
	value, ok := ca.Entries[key]
	return value.val, ok
}

func (ca *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			ca.mu.Lock()
			for key, val := range ca.Entries {
				timeElapsed := time.Since(val.createdAt)
				if timeElapsed > interval {
					delete(ca.Entries, key)
				}
			}
			ca.mu.Unlock()
		}
	}
}
