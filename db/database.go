package db

import (
	"fmt"
	"go-training/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	conn *gorm.DB
)

func Conn() *gorm.DB {
	return conn
}

func init() {
	// Construct the DSN (Data Source Name) for the PostgreSQL connection
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		config.Host, config.Port, config.Username, config.DatabaseName, config.Password)

	var err error
	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
}

func Ping() error {
	postgresDB, err := conn.DB()
	if err != nil {
		return err
	}

	return postgresDB.Ping()
}
