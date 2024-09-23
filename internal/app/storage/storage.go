package storage

import (
	"errors"
	"shortener/internal/app/models"
)

var (
	ErrStorageConnectionFailed = errors.New("storage db connection failed")
	ErrStorageFilePathEmpty    = errors.New("storage file path is empty")
	ErrStorageFileNotExists    = errors.New("storage file does not exist")
	ErrStorageValueNotFound    = errors.New("stored value not found")
)

type Storage interface {
	Connect() error
	Close() error
	Set(data models.ShortenData) error
	Get(ID string) (models.ShortenData, error)
}
