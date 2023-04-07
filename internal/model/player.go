package model

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name       string  `json:"name" gorm:"uniqueIndex:idx_player;not null;size:100"`
	SimpleName string  `json:"simple_name"`
	Sport      string  `json:"-" gorm:"not null"`
	Position   string  `json:"pos" gorm:"not null"`
	Age        float32 `json:"age" gorm:"not null"`
	TeamID     uint    `json:"team_id" gorm:"uniqueIndex:idx_player;not null"`
}

type PlayerAction struct {
	Player string
	Event  Event
}
