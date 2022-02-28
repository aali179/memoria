package models
import (
	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	ID           int    `json:"id"`
	Number       string `json:"number"`
	ImageOne     Image    `json:"imageOne" gorm:"foreignKey:ID"`
	ImageTwo     Image    `json:"imageTwo" gorm:"foreignKey:ID"`
    ImageThree   Image    `json:"imageThree" gorm:"foreignKey:ID"`
	HeadingOne   string `json:"headingOne"`
	HeadingTwo   string `json:"headingTwo"`
	HeadingThree string `json:"headingThree" `
	Song         Song   `json:"song" gorm:"foreignKey:ID"`
	Map         Map   `json:"map" gorm:"foreignKey:ID"`
}