package model

import "time"

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
	ID         uint       `json:"id" gorm:"primarykey"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"-" gorm:"index"`
	Event      Event      `gorm:"not null"`
	PlayerID   uint       `gorm:"not null"`
	GameID     uint       `gorm:"not null;index"`
	TS         float32    `gorm:"not null"`
	Quarter    int        `gorm:"not null"`
	PeriodType string     `gorm:"not null"`
	Period     int        `gorm:"not null"`
}

type GameLogResp struct {
	PlayerName string  `json:"player_name"`
	Event      Event   `json:"event"`
	TS         float32 `json:"time_stamp"`
	Quarter    int     `json:"quarter"`
	TeamName   string  `json:"team_name"`
	Score      uint    `json:"score,omitempty"`
}

type Action struct {
	TimeStamp string `json:"time_stamp"`
	Action    Event  `json:"action"`
}

type PlayerActionLog struct {
	Event map[int][]Action `json:"event"`
	Mins  map[int]float32  `json:"mins"`
	Pts   map[int]uint     `json:"pts"`
	Team  string           `json:"team"`
}

type PlayerData struct {
	PlayerName PlayerActionLog `json:"player_name"`
}

type PlayersData struct {
	PlayerData []PlayerData
}
