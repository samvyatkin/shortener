package storage

import (
	"context"
	"github.com/jackc/pgx/v5"
	"shortener/internal/app/models"
)

type DBStorage struct {
	connection string
}

func NewDBStorage(connection string) (*DBStorage, error) {
	if connection == "" {
		return nil, ErrStorageFilePathEmpty
	}

	return &DBStorage{
		connection: connection,
	}, nil
}

func (db *DBStorage) Connect() error {
	if db.connection == "" {
		return ErrStorageConnectionFailed
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, db.connection)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	return nil
}

func (db *DBStorage) Close() error {
	return nil
}

func (db *DBStorage) Set(data models.ShortenData) error {
	return nil
}

func (db *DBStorage) Get(ID string) (models.ShortenData, error) {
	return models.ShortenData{}, nil
}
