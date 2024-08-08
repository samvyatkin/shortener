package storage

import "errors"

var (
	ErrStorageValueNotFound = errors.New("storage value not found")
)

type Storage interface {
	Set(key string, value string)
	Get(key string) (string, error)
}
