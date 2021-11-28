package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	connection_string := "host=localhost port=5432 user=postgres dbname=challengehash sslmode=disable password=123"

	database, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection Error: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDatabase() *gorm.DB {
	return db
}
