package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Username     string
	Password     string
	DatabaseName string
	Host         string
	Port         string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables for database connection
	Username = os.Getenv("POSTGRES_USER")
	Password = os.Getenv("POSTGRES_PASSWORD")
	DatabaseName = os.Getenv("POSTGRES_DB")
	Host = os.Getenv("POSTGRES_HOST")
	Port = os.Getenv("POSTGRES_PORT")

	// Validate environment variables
	if Username == "" || Password == "" || DatabaseName == "" || Host == "" || Port == "" {
		log.Fatal("One or more required environment variables are missing.")
	}
}
