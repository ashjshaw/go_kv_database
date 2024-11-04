package store

import "sync"

type Store struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewStore() *Store {
	panic("NYI")
}

func (s *Store) Get(key string) ([]string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
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
	panic("NYI")
}
