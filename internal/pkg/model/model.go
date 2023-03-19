package model

import "time"

type commonModelFields struct {
	ID        uint       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" gorm:"index"`
}
