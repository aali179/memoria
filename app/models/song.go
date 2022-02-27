package models
import (
	"gorm.io/gorm"
)
type Song struct {
	gorm.Model
	ID        int
	SpotifyID string
}