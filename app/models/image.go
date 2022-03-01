package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	File []byte
	Map  []Map `json:"maps"`
}
