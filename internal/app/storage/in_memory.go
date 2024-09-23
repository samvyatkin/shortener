package storage

import (
	"shortener/internal/app/models"
	"sync"
)

type InMemoryStorage struct {
	cache map[string]models.ShortenData
	mutex sync.RWMutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		cache: make(map[string]models.ShortenData),
	}
}

func (s *InMemoryStorage) Connect() error {
	return nil
}

func (s *InMemoryStorage) Close() error {
	return nil
}

func (s *InMemoryStorage) Set(data models.ShortenData) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.cache[data.ID] = data
	return nil
}

func (s *InMemoryStorage) Get(ID string) (models.ShortenData, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	data, ok := s.cache[ID]
	if !ok {
		return models.ShortenData{}, ErrStorageValueNotFound
	}

	return data, nil
}
