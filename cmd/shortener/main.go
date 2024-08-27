package main

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"net/http"
	"shortener/internal/app/config"
	createRoute "shortener/internal/app/handlers/create"
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
	inMemoryStorage := storage.NewInMemoryStorage()

	// Handlers
	createHandler := createRoute.New(
		&utils.UUIDGenerator{},
		inMemoryStorage,
		c,
	)
	searchHandler := search.New(inMemoryStorage)
	shortenHandler := shorten.New(
		&utils.UUIDGenerator{},
		inMemoryStorage,
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
		http.ListenAndServe(c.ServerAddr(), middleware.WithLogging(utils.Logger)(r)),
	)
}
