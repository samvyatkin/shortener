package main

import (
	"bufio"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
	"shortener/internal/app/config"
	createRoute "shortener/internal/app/handlers/create"
	"shortener/internal/app/handlers/search"
	"shortener/internal/app/handlers/shorten"
	"shortener/internal/app/middleware"
	"shortener/internal/app/models"
	"shortener/internal/app/storage"
	"shortener/internal/app/utils"
	"time"
)

func main() {
	run()
}

func run() {
	defer utils.Logger.Sync()

	// Configuration
	c := config.NewConfig()
	r := chi.NewRouter()

	// Middleware
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.Timeout(60 * time.Second))

	// Stored data
	data, err := Load(c.FileStoragePath())
	if err != nil {
		utils.Logger.Fatal("failed to load stored shorten urls", err)
	}

	// Services
	inMemoryStorage := storage.NewInMemoryStorage(data)
	fileStorage, err := storage.NewFileStorage(c.FileStoragePath())
	if err != nil {
		utils.Logger.Fatal("failed to init storage", err)
	}

	// Handlers
	createHandler := createRoute.New(
		&utils.UUIDGenerator{},
		inMemoryStorage,
		fileStorage,
		c,
	)
	searchHandler := search.New(inMemoryStorage)
	shortenHandler := shorten.New(
		&utils.UUIDGenerator{},
		inMemoryStorage,
		fileStorage,
		c,
	)

	// Routes
	r.Route("/", func(r chi.Router) {
		r.Post("/", createHandler.Handle)
		r.Post("/api/shorten", shortenHandler.Handle)
		r.Get("/{id}", searchHandler.Handle)
	})

	utils.Logger.Infow("Running server", "addr", c.ShortenerAddr())
	utils.Logger.Fatalw(
		"Can't run server",
		http.ListenAndServe(
			c.ServerAddr(),
			middleware.WithLogging(utils.Logger)(
				middleware.WithCompress()(r),
			),
		),
	)
}

func Load(path string) ([]models.ShortenData, error) {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data := make([]models.ShortenData, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		shortenData := models.ShortenData{}

		err := json.Unmarshal([]byte(line), &shortenData)
		if err != nil {
			return nil, err
		}

		data = append(data, shortenData)
	}

	if scanner.Err() != nil {
		return nil, err
	}

	return data, nil
}
