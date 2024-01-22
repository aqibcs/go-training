// Package db provides functionality for establishing and managing a PostgreSQL database connection using GORM.
package db

import (
	"fmt"
	"go-training/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// conn is the global database connection.
	conn *gorm.DB
)

// Conn returns the global database connection.
func Conn() *gorm.DB {
	return conn
}

// init initializes the database connection during package initialization.
func init() {
	// Construct the Data Source Name (DSN) for the PostgreSQL connection.
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		config.Host, config.Port, config.Username, config.DatabaseName, config.Password)

	// Initialize the GORM database connection.
	var err error
	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to the database: %v", err)
		// Handle the error gracefully, don't exit the application.
	}
}

// Ping checks the health of the database connection by attempting to ping the underlying database.
// It returns an error if the ping operation fails.
func Ping() error {
	postgresDB, err := conn.DB()
	if err != nil {
		return err
	}

	return postgresDB.Ping()
}
