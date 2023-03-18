package model

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Date   time.Time `gorm:"uniqueIndex:idx_game;not null"`
	AwayID uint
	HomeID uint `gorm:"uniqueIndex:idx_game;not null"`
}
