package models

import (
	"gorm.io/gorm"
)

type Map struct {
	gorm.Model
	Location string
	ImageID  uint `json:"image"`
	PageID   uint `json:"page"`
}
