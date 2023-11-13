package main

import (
	"github.com/labstack/echo/v4"
	"go-training/auth"
	"go-training/db"
	"go-training/handlers"
	models "go-training/models/object"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// public routes
	e.GET("/", handlers.GetJwt)

	// protected routes
	apiGroup := e.Group("/api")
	apiGroup.Use(auth.ValidateJWT)
	apiGroup.GET("/employee", handlers.GetAllEmployees)
	apiGroup.GET("/employee/:employee_id", handlers.GetEmployeeByID)
	apiGroup.POST("/employee", handlers.CreateEmployee)
	apiGroup.PATCH("/employee/:employee_id", handlers.UpdateEmployee)
	apiGroup.DELETE("/employee/:employee_id", handlers.DeleteEmployee)
	apiGroup.GET("/upload", handlers.UploadFileHandler)
	apiGroup.POST("/hello", handlers.HelloHandler)

	// Auto migrate database
	dbConn := db.Conn()
	dbConn.AutoMigrate(&models.Employee{})

	// Start the Echo server on port 8080
	e.Start(":8080")
}