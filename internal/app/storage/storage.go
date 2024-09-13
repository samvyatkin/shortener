package storage

import (
	"errors"
	"shortener/internal/app/models"
)

var (
	ErrStorageFilePathEmpty = errors.New("storage file path is empty")
	ErrStorageFileNotExists = errors.New("storage file does not exist")
	ErrStorageValueNotFound = errors.New("stored value not found")
)

type Storage interface {
	Set(data models.ShortenData) error
	Get(ID string) (models.ShortenData, error)
}
