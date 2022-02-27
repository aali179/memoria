package models
import (
	"gorm.io/gorm"
)
type Image struct {
	gorm.Model
	ID int
	File string
}
