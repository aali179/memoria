package models

import (
	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	Number       string  `json:"number"`
	Images       []Image `gorm:"many2many:page_images;"`
	HeadingOne   string  `json:"headingOne"`
	HeadingTwo   string  `json:"headingTwo"`
	HeadingThree string  `json:"headingThree" `
	Song         Song    `json:"song"`
	Map          Map     `json:"map"`
	ScrapbookID  uint    `json:"scrapbook"`
}
