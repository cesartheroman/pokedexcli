package main

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: ttl,
	}
	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	fmt.Printf("Adding url: %s to cache\n", key)
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, ok := c.entries[key]; ok {
		fmt.Println("Entry found in cache for url:", key)
		return entry.value, ok
	} else {
		return []byte{}, ok
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()

		c.mu.Lock()
		for key, entry := range c.entries {
			if now.Sub(entry.createdAt) > c.interval {
				fmt.Println("Cache expired!")
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
