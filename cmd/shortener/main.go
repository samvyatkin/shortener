package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"shortener/internal/app/config"
	createRoute "shortener/internal/app/handlers/create"
	"shortener/internal/app/handlers/search"
	"shortener/internal/app/storage"
	"shortener/internal/app/utils"
	"time"
)

func main() {
	run()
}

func run() {
	// Configuration
	c := config.NewConfig()
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Timeout(60 * time.Second))

	// Services
	inMemoryStorage := storage.NewInMemoryStorage()

	// Handlers
	createHandler := createRoute.New(
		&utils.UUIDGenerator{},
		inMemoryStorage,
		c,
	)
	searchHandler := search.New(inMemoryStorage)

	// Routes
	r.Route("/", func(r chi.Router) {
		r.Post("/", createHandler.Handle)
		r.Get("/{id}", searchHandler.Handle)
	})

	log.Fatal(http.ListenAndServe(c.ServerAddr(), r))
}
