package main

import (
	"net/http"
	createRoute "shortener/internal/app/handlers/shortener_create_post"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	createHandler := createRoute.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/", createHandler.Handle)

	return http.ListenAndServe(`:8080`, mux)
}
