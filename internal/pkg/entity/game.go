package entity

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	AwayTeam          string
	HomeTeam          string
	AwayTeamPlayerIDs string
	HomeTeamPlayerIDS string
	Date              time.Time
}
