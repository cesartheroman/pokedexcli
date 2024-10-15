package main

import (
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

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		value:     val,
	}

	c.entries[key] = entry
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		entries:  make(map[string]cacheEntry),
		interval: ttl,
	}
}
