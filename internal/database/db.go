package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Client() *gorm.DB {
	return db
}

func Connect() {
	dsn:=os.Getenv("DB_ENV")

	if dsn==""{
		log.Fatalf("unable to fetch env variables")
	}

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to connect to database : %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("unable to connect to db: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("error while pinging db : %v", err)
	}
}
