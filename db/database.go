package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

	// Absolute path to schema.sql file
	filePath := "/home/aqib/work/go-training/db/schema.sql"

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("Error: schema.sql file not found at %s", filePath)
	}
}

func ApplySchemaFromFile(filePath string) {
	// Read SQL statements from schema.sql file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	sqlStatements := string(content)

	// Split SQL statements and execute them
	statements := strings.Split(sqlStatements, ";")
	for _, statement := range statements {
		statement = strings.TrimSpace(statement)
		if statement != "" {
			err := DB.Exec(statement).Error
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
