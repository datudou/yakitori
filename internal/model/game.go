package model

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Season        string    `json:"season" gorm:"index:,unique,composite:idx_game"`
	StartDate     time.Time `json:"date" gorm:"index:,unique,composite:idx_game"`
	AwayTeam      string    `json:"away_team" gorm:"not null"`
	HomeTeam      string    `json:"home_team" gorm:"index:,unique,composite:idx_game"`
	AwayTeamScore int       `json:"away_score" gorm:"not null"`
	HomeTeamScore int       `json:"home_score" gorm:"not null"`
}

type GameResp struct {
	Game
	AwayTeamIcon string `json:"away_team_icon"`
	HomeTeamIcon string `json:"home_team_icon"`
}
