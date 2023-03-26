package model

import (
	"time"
)

type Game struct {
	ID            uint       `json:"id" gorm:"primarykey"`
	CreatedAt     time.Time  `json:"-"`
	UpdatedAt     time.Time  `json:"-"`
	DeletedAt     *time.Time `json:"-" gorm:"index"`
	Season        string     `json:"season" gorm:"index:,unique,composite:idx_game"`
	StartDate     time.Time  `json:"date" gorm:"index:,unique,composite:idx_game"`
	AwayTeam      string     `json:"away_team" gorm:"not null"`
	HomeTeam      string     `json:"home_team" gorm:"index:,unique,composite:idx_game"`
	AwayTeamScore int        `json:"away_score" gorm:"not null"`
	HomeTeamScore int        `json:"home_score" gorm:"not null"`
}

type GameResp struct {
	Game
	AwayTeamIcon string `json:"away_team_icon"`
	HomeTeamIcon string `json:"home_team_icon"`
}
