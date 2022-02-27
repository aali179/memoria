package models
import (
	"gorm.io/gorm"
)
type Map struct {
	gorm.Model
	Location    string
	Image     Image    `json:"image" gorm:"foreignKey:ID"`
}
