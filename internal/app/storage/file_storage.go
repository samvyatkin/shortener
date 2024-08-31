package storage

import (
	"encoding/json"
	"os"
	"shortener/internal/app/models"
	"sync"
)

type FileStorage struct {
	file  *os.File
	mutex sync.Mutex
}

func NewFileStorage(path string) (*FileStorage, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &FileStorage{
		file: file,
	}, nil
}

func (fs *FileStorage) Close() error {
	return fs.file.Close()
}

func (fs *FileStorage) Set(data models.ShortenData) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	d = append(d, '\n')

	_, err = fs.file.Write(d)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileStorage) Get(ID string) (models.ShortenData, error) {
	return models.ShortenData{}, nil
}
