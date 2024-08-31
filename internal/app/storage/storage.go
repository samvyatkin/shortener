package storage

import (
	"errors"
	"shortener/internal/app/models"
)

var (
	ErrStorageValueNotFound = errors.New("stored value not found")
)

type Storage interface {
	Set(data models.ShortenData) error
	Get(ID string) (models.ShortenData, error)
}
