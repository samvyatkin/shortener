package storage

type InMemoryStorage struct {
	cache map[string]string
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		cache: make(map[string]string),
	}
}

func (s *InMemoryStorage) Set(key, value string) {
	s.cache[key] = value
}

func (s *InMemoryStorage) Get(key string) (string, error) {
	if val, ok := s.cache[key]; ok {
		return val, nil
	}

	return "", ErrStorageValueNotFound
}
