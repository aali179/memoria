package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	SpotifyID string `json:"spotifyID"`
	PageID    uint   `json:"page"`
}
