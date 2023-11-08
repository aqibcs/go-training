package db

import (
	"go-training/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	conn *gorm.DB
)

func Conn() *gorm.DB {
	return conn
}

func init() {
	// Load environment variables from .env file using the config package
	config.LoadEnv()

	// Construct the DSN (Data Source Name) for the PostgreSQL connection
	dsn := "user=" + config.Username + " password=" + config.Password + " dbname=" + config.DatabaseName + " host=" + config.Host + " port=" + config.Port

	var err error
	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
}

func Ping() error {
	postgresDB, err := conn.DB()
	if err != nil {
		return err
	}

	return postgresDB.Ping()
}
