package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"go-training/db"
	"go-training/handlers"
)

func main() {
	// Initialize the database connection
	db.Init()

	// Create a new Chi router
	r := chi.NewRouter()

	// Routes and corresponding handlers
	r.Get("/upload", handlers.UploadFileHandler)
	r.Post("/hello", handlers.HelloHandler)

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", r)
}
