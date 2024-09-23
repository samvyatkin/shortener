package main

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"net/http"
	"shortener/internal/app/config"
	createRoute "shortener/internal/app/handlers/create"
	"shortener/internal/app/handlers/ping"
	"shortener/internal/app/handlers/search"
	"shortener/internal/app/handlers/shorten"
	"shortener/internal/app/middleware"
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

	// Services
	s, err := URLsStorage(c)
	if err != nil {
		utils.Logger.Fatalf("failed to init storage %v", err)
	}

	// Handlers
	createHandler := createRoute.New(
		&utils.UUIDGenerator{},
		s,
		c,
	)
	searchHandler := search.New(s)
	shortenHandler := shorten.New(
		&utils.UUIDGenerator{},
		s,
		c,
	)
	pingHandler := ping.New(
		s,
		c,
	)

	// Routes
	r.Route("/", func(r chi.Router) {
		r.Post("/", createHandler.Handle)
		r.Post("/api/shorten", shortenHandler.Handle)
		r.Get("/{id}", searchHandler.Handle)
		r.Get("/ping", pingHandler.Handle)
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

func URLsStorage(config *config.Config) (storage.Storage, error) {
	dbStorage, err := storage.NewDBStorage(config.DBAddr())
	if err != nil {
		fileStorage, err := storage.NewFileStorage(config.FileStoragePath())
		if err != nil {
			return storage.NewInMemoryStorage(), nil
		}

		return fileStorage, nil
	}

	return dbStorage, nil
}
