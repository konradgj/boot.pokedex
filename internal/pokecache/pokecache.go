package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.RWMutex
	stop    chan struct{}
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: map[string]cacheEntry{},
		stop:    make(chan struct{}),
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	c.entries[key] = newEntry

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) Stop() {
	close(c.stop)
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			c.mu.Lock()
			for key, entry := range c.entries {
				if entry.createdAt.Add(interval).Before(t) {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		case <-c.stop:
			return
		}
	}
}
