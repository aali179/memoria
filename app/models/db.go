package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {

	godotenv.Load()
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, username, name, password)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})

	if err != nil {
		fmt.Println("err")
	}

	db = conn
	db.Debug().AutoMigrate(&Scrapbook{}, &Page{}, &Image{}, &Song{}, &Map{})
}

func GetPageById(id uint) Page {
	var page Page
	db.Where("id = ?", id).First(&page)
	return page
}

func CreatePage(
	imageOne Image,
	imageTwo Image,
	imageThree Image,
	headingOne string,
	headingTwo string,
	headingThree string,
	location Map,
	song Song,
	scrapbookID uint) Page {
	page := Page{Images: []Image{imageOne, imageTwo, imageThree}, HeadingOne: headingOne, HeadingTwo: headingTwo, HeadingThree: headingThree, Map: location, Song: song, ScrapbookID: scrapbookID}
	db.Create(&page)
	return page
}
