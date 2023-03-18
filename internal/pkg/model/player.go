package model

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name          string `gorm:"uniqueIndex:idx_player;not null;size:100"`
	Sport         string
	Position      string
	UsedTeamIDs   string
	PlayerAvatar  string
	Age           int
	CurrentTeamID uint `gorm:"uniqueIndex:idx_player;not null"`
}
