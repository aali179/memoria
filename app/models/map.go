package models

import (
	"gorm.io/gorm"
)

type Map struct {
	gorm.Model
	Location string `json:"location"`
	ImageID  uint   `json:"image"`
	PageID   uint   `json:"page"`
}
