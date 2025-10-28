package simplecache

import (
	"fmt"
	"sync"
)

type ConcurrentStore struct {
	mu    sync.Mutex
	Store map[string]any
}

func (s *ConcurrentStore) Set(key string, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Store[key] = value
}

func (s *ConcurrentStore) Get(key string) (any, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if key == "" {
		fmt.Println("Empty key")
	}
	value, exists := s.Store[key]
	return value, exists
}
