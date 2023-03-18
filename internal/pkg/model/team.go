package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Code       string `json:"code" gorm:"index"`
	Name       string `json:"name" gorm:"uniqueIndex:idx_team_name;not null" `
	League     string `json:"league"`
	SimpleName string `json:"simple_name"`
	TeamLogo   string `json:"team_logo"`
	Conference string `json:"conference"`
	Division   string `json:"division"`
	Category   string `json:"category"`
}
