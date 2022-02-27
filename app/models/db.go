
package models

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"os"
)

var db *gorm.DB

func Connect() {

	godotenv.Load()
	
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, username, name, password)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	if err != nil {
		fmt.Println("err")
	}

	db = conn
	db.Debug().AutoMigrate(&Map{}, &Song{}, &Image{},&Scrapbook{},&Page{})
}
