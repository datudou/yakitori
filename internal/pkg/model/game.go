package model

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Season        string    `gorm:"index:,unique,composite:idx_game"`
	Date          time.Time `gorm:"index:,unique,composite:idx_game"`
	AwayTeam      string    `gorm:"not null"`
	HomeTeam      string    `gorm:"index:,unique,composite:idx_game"`
	AwayTeamScore int       `gorm:"not null"`
	HomeTeamScore int       `gorm:"not null"`
}
