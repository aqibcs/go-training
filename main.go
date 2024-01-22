package main

import (
	"go-training/auth"
	"go-training/db"
	"go-training/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-training/models/object"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Middleware for recovering from panics
	e.Use(middleware.Recover())

	// Public routes
	e.GET("/get-jwt", handlers.GetJwt)

	// Protected routes
	apiGroup := e.Group("/api")
	apiGroup.Use(auth.ValidateJWT)

	// Routes and corresponding handlers under the "/api" group
	apiGroup.GET("/employee", handlers.GetAllEmployees)
	apiGroup.GET("/employee/:employee_id", handlers.GetEmployeeByID)
	apiGroup.POST("/employee", handlers.CreateEmployee)
	apiGroup.PATCH("/employee/:employee_id", handlers.UpdateEmployee)
	apiGroup.DELETE("/employee/:employee_id", handlers.DeleteEmployee)
	apiGroup.GET("/upload", handlers.UploadFileHandler)
	apiGroup.POST("/hello", handlers.HelloHandler)

	// Migrate database
	dbConn := db.Conn()
	dbConn.AutoMigrate(&models.Employee{})

	// Start the HTTP server on port 8080
	e.Start(":8080")
}
