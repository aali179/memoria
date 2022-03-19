package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	File string
	Map  []Map `json:"maps"`
}
