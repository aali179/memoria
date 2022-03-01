package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() error {

	godotenv.Load()
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, username, name, password)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	if err != nil {
		return err
	}

	db = conn
	//db.Debug().AutoMigrate(&Scrapbook{}, &Page{}, &Image{}, &Song{}, &Map{})

	return nil
}

func GetPageById(id uint) (Page, error) {
	var page Page
	res := db.Where("id = ?", id).First(&page)
	if res.Error != nil {
		return Page{}, res.Error
	}
	return page, nil
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
	scrapbookID uint) (Page, error) {
	page := Page{Images: []Image{imageOne, imageTwo, imageThree}, HeadingOne: headingOne, HeadingTwo: headingTwo, HeadingThree: headingThree, Map: location, Song: song, ScrapbookID: scrapbookID}
	res := db.Create(&page)
	if res.Error != nil {
		return Page{}, res.Error
	}
	return page, nil
}

func CreateScrapbook(name string) (Scrapbook, error) {
	scrapbook := Scrapbook{Name: name}
	res := db.Create(&scrapbook)
	if res.Error != nil {
		return Scrapbook{}, res.Error
	}
	return scrapbook, nil
}
