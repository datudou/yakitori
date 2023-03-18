package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name     string  `gorm:"uniqueIndex:idx_player;not null;size:100"`
	Sport    string  `gorm:"not null"`
	Position string  `gorm:"not null"`
	Age      float32 `gorm:"not null"`
	TeamID   uint    `gorm:"uniqueIndex:idx_player;not null"`
}
