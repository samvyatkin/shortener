package mocks

import "shortener/internal/app/storage"

type InMemoryStorageMock struct {
	cache map[string]string
}

func NewInMemoryStorageMock(data map[string]string) *InMemoryStorageMock {
	return &InMemoryStorageMock{
		cache: data,
	}
}

func (s *InMemoryStorageMock) Set(key string, value string) {
	s.cache[key] = value
}

func (s *InMemoryStorageMock) Get(key string) (string, error) {
	if v, ok := s.cache[key]; ok {
		return v, nil
	}

	return "", storage.ErrStorageValueNotFound
}
