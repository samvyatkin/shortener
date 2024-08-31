package mocks

import (
	"shortener/internal/app/models"
	"shortener/internal/app/storage"
)

type InMemoryStorageMock struct {
	cache map[string]models.ShortenData
}

func NewInMemoryStorageMock(data map[string]models.ShortenData) *InMemoryStorageMock {
	return &InMemoryStorageMock{
		cache: data,
	}
}

func (s *InMemoryStorageMock) Set(data models.ShortenData) error {
	s.cache[data.ID] = data
	return nil
}

func (s *InMemoryStorageMock) Get(ID string) (models.ShortenData, error) {
	if d, ok := s.cache[ID]; ok {
		return d, nil
	}

	return models.ShortenData{}, storage.ErrStorageValueNotFound
}
