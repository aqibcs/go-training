package main

import (
	"github.com/go-chi/chi"
	"go-training/handlers"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/upload", handlers.UploadFileHandler)
	r.Post("/hello", handlers.HelloHandler)
	http.ListenAndServe(":8080", r)
}
