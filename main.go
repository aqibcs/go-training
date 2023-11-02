package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"go-training/db"
	"go-training/handlers"
	"go-training/models/object"
)

func main() {
	// Initialize the database connection
	db.Init()

	// Create a new Chi router
	r := chi.NewRouter()

	// Routes and corresponding handlers
	r.Get("/object", handlers.GetAllObjects)
	r.Get("/object/{object_id}", handlers.GetObjectByID)
	r.Post("/object", handlers.CreateObject)
	r.Patch("/object/{object_id}", handlers.UpdateObject)
	r.Delete("/object/{object_id}", handlers.DeleteObject)
	r.Get("/upload", handlers.UploadFileHandler)
	r.Post("/hello", handlers.HelloHandler)

	dbConn := db.GetDB()
	dbConn.AutoMigrate(&models.Object{})

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", r)
}
