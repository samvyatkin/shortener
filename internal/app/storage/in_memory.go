package storage

import "sync"

type InMemoryStorage struct {
	cache sync.Map
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		cache: sync.Map{},
	}
}

func (s *InMemoryStorage) Set(key, value string) {
	s.cache.Store(key, value)
}

func (s *InMemoryStorage) Get(key string) (string, error) {
	v, ok := s.cache.Load(key)
	if !ok {
		return "", ErrStorageValueNotFound
	}

	return v.(string), nil
}
