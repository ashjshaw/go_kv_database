package store

import (
	"maps"
	"slices"
	"sync"
)

type Store struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewStore() *Store {
	return &Store{
		data: map[string]string{},
		mu:   sync.RWMutex{},
	}
}

func (s *Store) Get(key string) ([]string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if key == "" {
		return slices.Collect(maps.Keys(s.data)), true
	}
	if value, ok := s.data[key]; ok {
		return []string{value}, true
	}
	return nil, false
}

func (s *Store) Put(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *Store) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.data[key]
	if ok {
		delete(s.data, key)
	}
	return ok
}
