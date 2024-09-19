package storage

import (
	"bufio"
	"encoding/json"
	"os"
	"shortener/internal/app/models"
	"shortener/internal/app/utils"
	"strings"
	"sync"
)

type FileStorage struct {
	path  string
	mutex *sync.RWMutex
}

func NewFileStorage(path string) (*FileStorage, error) {
	if path == "" {
		return nil, ErrStorageFilePathEmpty
	}

	return &FileStorage{
		path:  path,
		mutex: &sync.RWMutex{},
	}, nil
}

func (fs *FileStorage) Set(data models.ShortenData) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	file, err := os.OpenFile(fs.path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	d = append(d, '\n')

	_, err = file.Write(d)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileStorage) Get(ID string) (models.ShortenData, error) {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()

	file, err := os.OpenFile(fs.path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		utils.Logger.Errorw("Can't open file %s", fs.path)
		return models.ShortenData{}, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		shortenData := models.ShortenData{}

		err := json.Unmarshal([]byte(line), &shortenData)
		if err != nil {
			utils.Logger.Errorw("Can't parse object: text - %s err - %w", line, err)
			return models.ShortenData{}, err
		}

		if shortenData.ID == ID {
			return shortenData, nil
		}
	}

	if err := scanner.Err(); err != nil {
		utils.Logger.Errorw("Scanner error %w", err)
		return models.ShortenData{}, err
	}

	return models.ShortenData{}, ErrStorageValueNotFound
}
