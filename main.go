package main

import (
	"go-training/auth"
	"go-training/db"
	"go-training/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"go-training/models/object"
)

func main() {
	// Create a new Chi router
	r := chi.NewRouter()

	// public routes
	r.Group(func(r chi.Router) {
		r.Get("/", handlers.GetJwt)
	})

	// protected routes
	r.Group(func(r chi.Router) {
		r.Use(auth.ValidateJWT)
		r.Route("/api", func(r chi.Router) {
			// Routes and corresponding handlers under the "/api" group
			r.Get("/object", handlers.GetAllObjects)
			r.Get("/object/{object_id}", handlers.GetObjectByID)
			r.Post("/object", handlers.CreateObject)
			r.Patch("/object/{object_id}", handlers.UpdateObject)
			r.Delete("/object/{object_id}", handlers.DeleteObject)
			r.Get("/upload", handlers.UploadFileHandler)
			r.Post("/hello", handlers.HelloHandler)
		})
	})

	dbConn := db.Conn()
	dbConn.AutoMigrate(&models.Employee{})

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", r)
}
