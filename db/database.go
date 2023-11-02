package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	conn    *gorm.DB
	dbMutex sync.Mutex
)

func GetDB() *gorm.DB {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	return conn
}

// var DB *gorm.DB

func Init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables for database connection
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	// Construct the DSN (Data Source Name) for the PostgreSQL connection
	dsn := "user=" + username + " password=" + password + " dbname=" + databaseName + " host=" + host + " port=" + port

	var errDB error
	conn, errDB = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatal("Failed to connect to the database")
	}
}

func PingDB() error {
	postgresDB, err := conn.DB()
	if err != nil {
		return err
	}
	err = postgresDB.Ping()
	if err != nil {
		return err
	}
	return nil
}
