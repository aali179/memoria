package models
import (
	"gorm.io/gorm"
)

type Scrapbook struct {
	gorm.Model
	Page []Page `json:"page" gorm:"foreignKey:ID"`
	Name string `json:"name"`
}

