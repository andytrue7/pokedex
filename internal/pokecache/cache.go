package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    sync.Mutex
	cache map[string]cacheEntry
}

type cacheEntry struct {
	value     []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		value:     value,
		createdAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[key]
	return entry.value, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.cache {
		if time.Since(entry.createdAt) > interval {
			delete(c.cache, key)
		}
	}
}
