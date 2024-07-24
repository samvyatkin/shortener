package main

import (
	"net/http"
	createRoute "shortener/internal/app/handlers/create"
	"shortener/internal/app/storage"
	"shortener/internal/app/utils"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	createHandler := createRoute.New(
		&utils.UUIDGenerator{},
		storage.NewInMemoryStorage(),
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/", createHandler.Handle)

	return http.ListenAndServe(`:8080`, mux)
}
