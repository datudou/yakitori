package entity

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name          string
	Sport         string
	Position      string
	CurrentTeamID uint
	UsedTeamIDs   string
	PlayerAvatar  string
	Age           int
}
