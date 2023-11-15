package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Username     = "myuser"
	Password     = ""
	DatabaseName = "mydatabase"
	Host         = "localhost"
	Port         = "5432"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
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
