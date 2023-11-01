package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	// postgres://<user_name>:<password>@localhost:<port>/<database_name>
	dsn := "postgres://myuser:mypassword@localhost:5432/mydatabase"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Ensure the database connection is valid
	postgresDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	err = postgresDB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
