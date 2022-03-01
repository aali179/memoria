package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	SpotifyID string
	PageID    uint `json:"page"`
}
