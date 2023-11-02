package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	dB      *gorm.DB
	dbMutex sync.Mutex
)

func GetDB() *gorm.DB {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	return dB
}

func Init() {
	// postgres://<user_name>:<password>@localhost:<port>/<database_name>
	dsn := "postgres://myuser:mypassword@localhost:5432/mydatabase"
	var err error
	dB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
}

func PingDB() error {
	postgresDB, err := dB.DB()
	if err != nil {
		return err
	}
	err = postgresDB.Ping()
	if err != nil {
		return err
	}
	return nil
}
