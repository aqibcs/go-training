package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	conn *gorm.DB
)

func GetDB() *gorm.DB {
	return conn
}

func Init() {
	// postgres://<user_name>:<password>@localhost:<port>/<database_name>
	dsn := "postgres://myuser:mypassword@localhost:5432/mydatabase"
	var err error
	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
}

func PingDB() error {
	postgresDB, err := conn.DB()
	if err != nil {
		return err
	}

	return postgresDB.Ping()
}
