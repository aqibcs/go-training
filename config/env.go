package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// PostgreSQL connection parameters
var (
	Username     = "myuser"
	Password     = ""
	DatabaseName = "mydatabase"
	Host         = "localhost"
	Port         = "5432"
)

// init loads environment variables from a .env file and overrides default values.
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	if x := os.Getenv("POSTGRES_USER"); x != "" {
		Username = x
	}

	if x := os.Getenv("POSTGRES_PASSWORD"); x != "" {
		Password = x
	}

	if x := os.Getenv("POSTGRES_DB"); x != "" {
		DatabaseName = x
	}

	if x := os.Getenv("POSTGRES_HOST"); x != "" {
		Host = x
	}

	if x := os.Getenv("POSTGRES_PORT"); x != "" {
		Port = x
	}
}
