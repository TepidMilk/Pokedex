package pokeapi

import (
	"sync"
	"time"
)

type Cache struct {
	Entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) *Cache {
	newCache := &Cache{
		Entries:  make(map[string]cacheEntry),
		interval: duration,
	}
	go newCache.reapLoop()
	return newCache
}

func (ca Cache) Add(key string, val []byte) {
	ca.mu.Lock()
	defer ca.mu.Unlock()

	ca.Entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (ca Cache) Get(key string) ([]byte, bool) {
	ca.mu.Lock()
	defer ca.mu.Unlock()
	value, ok := ca.Entries[key]
	if ok {
		return value.val, true
	}
	return nil, false
}

func (ca Cache) reapLoop() {
	ticker := time.NewTicker(ca.interval)
	for {
		select {
		case <-ticker.C:
			ca.mu.Lock()
			for key, val := range ca.Entries {
				timeElapsed := time.Since(val.createdAt)
				if timeElapsed > ca.interval {
					delete(ca.Entries, key)
				}
			}
			ca.mu.Unlock()
		}
	}
}
