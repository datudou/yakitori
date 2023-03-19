package model

import "time"

type Player struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" gorm:"index"`
	Name      string     `json:"name" gorm:"index:,unique,composite:idx_player"`
	Sport     string     `json:"-" gorm:"not null"`
	Position  string     `json:"pos" gorm:"not null"`
	Age       float32    `json:"age" gorm:"not null"`
	TeamID    uint       `json:"team_id" gorm:"index:,unique,composite:idx_player"`
}
