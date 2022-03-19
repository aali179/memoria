package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbTest *gorm.DB

func ConnectTestDB() error {
	godotenv.Load()
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	name := os.Getenv("TEST_DB_NAME")
	host := os.Getenv("DB_HOST")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, username, name, password)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	if err != nil {
		return err
	}

	db = conn

	return nil
}

func GetTestDB() *gorm.DB {
	return dbTest
}
