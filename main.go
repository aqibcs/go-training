package main

import (
	"go-training/handlers"
	"net/http"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
    r.Get("/upload", handlers.UploadFileHandler)
	r.Post("/hello", handlers.HelloHandler)
	http.ListenAndServe(":8080", r)
}
