package simplecache

import (
	"fmt"
	"sync"
)

type ConcurrentStore struct {
	mu    sync.RWMutex
	Store map[string]any
}

type Storer interface {
	Set(key string, value any)
	Get(key string) (any, bool)
}

func (c *ConcurrentStore) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if key != "" && value != nil {
		c.Store[key] = value
	}
}

func (c *ConcurrentStore) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if key == "" {
		fmt.Printf("The key is missing: %q\n", key)
	}
	value, exists := c.Store[key]

	return value, exists
}
