package model

import "time"

type Team struct {
	ID         uint       `json:"id" gorm:"primarykey"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"-" gorm:"index"`
	Code       string     `json:"code" gorm:"index"`
	Name       string     `json:"name" gorm:"uniqueIndex:idx_team_name;not null;size:100" `
	League     string     `json:"league"`
	SimpleName string     `json:"simple_name"`
	TeamLogo   string     `json:"team_logo"`
	Conference string     `json:"conference"`
	Division   string     `json:"division"`
	Category   string     `json:"category"`
}
