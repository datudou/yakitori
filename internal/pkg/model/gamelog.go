package model

import (
	"gorm.io/gorm"
)

type Event string

const (
	BLOCK      Event = "block"
	ASSIST     Event = "assist"
	MAKE3      Event = "make-3"
	MAKE2      Event = "make-2"
	MAKEFT     Event = "make-ft"
	MISS3      Event = "miss-3"
	MISS2      Event = "miss-2"
	MISSFT     Event = "miss-ft"
	TO         Event = "turnover"
	STEAL      Event = "steal"
	DEFREBOUND Event = "def-reb"
	OFFREBOUND Event = "off-reb"
)

type GameLog struct {
	gorm.Model
	Event    Event   `gorm:"not null"`
	PlayerID uint    `gorm:"not null"`
	GameID   uint    `gorm:"not null;index"`
	TS       float32 `gorm:"not null"`
	Quarter  int     `gorm:"not null"`
}
