package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name     string  `gorm:"index:,unique,composite:idx_player"`
	Sport    string  `gorm:"not null"`
	Position string  `gorm:"not null"`
	Age      float32 `gorm:"not null"`
	TeamID   uint    `gorm:"index:,unique,composite:idx_player"`
}
