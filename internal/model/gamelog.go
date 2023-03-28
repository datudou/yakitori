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
	ON         Event = "on"
	OFF        Event = "off"
)

type GameLog struct {
	ID                       uint       `json:"id" gorm:"primarykey"`
	CreatedAt                time.Time  `json:"-"`
	UpdatedAt                time.Time  `json:"-"`
	DeletedAt                *time.Time `json:"-" gorm:"index"`
	PlayerID                 uint       `gorm:"index;not null"`
	GameID                   uint       `gorm:"index;not null"`
	RemainingSecondsInPeriod float32    `gorm:"not null"`
	Period                   int        `gorm:"not null"`
	Event                    Event      `gorm:"not null"`
	PeriodType               string     `gorm:"not null"`
	AwayScore                int        `gorm:"not null"`
	HomeScore                int        `gorm:"not null"`
	Description              string     `gorm:"not null"`
	EventID                  string     `gorm:"uniqueIndex:idx_gamelog;not null;size:100"`
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
	TimeStamp float32 `json:"time_stamp"`
	Action    Event   `json:"action"`
}

type PlayerActionLog struct {
	Event     map[int][]Action            `json:"event"`
	Mins      map[int][]map[Event]float32 `json:"mins"`
	Team      string                      `json:"team"`
	IsOnField int                         `json:"-"`
	Period    int                         `json:"-"`
}

type PlayerData struct {
	PlayerName PlayerActionLog `json:"player_name"`
}

type PlayersData struct {
	PlayerData []PlayerData
}
