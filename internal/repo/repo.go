package repo

import (
	"context"
	"time"

	"github.com/new-pop-corn/internal/model"
	"gorm.io/gorm"
)

type IPlayerRepo interface {
	GetPlayersByTeamID(ctx context.Context, teamID uint) (*model.Player, error)
	GetPlayerByName(ctx context.Context, playerName string) (*model.Player, error)
	GetPlayers(ctx context.Context, league string) ([]*model.Player, error)
	CreatePlayer(ctx context.Context, player model.Player) (*model.Player, error)
	GetPlayerByID(ctx context.Context, id uint) (*model.Player, error)
	GetPlayerBySimpleName(ctx context.Context, playerName string, teamID uint) (*model.Player, error)
}
type ITeamRepo interface {
	GetTeamByID(ctx context.Context, teamID uint) (*model.Team, error)
	GetTeamByName(ctx context.Context, name string) (*model.Team, error)
	GetTeamByCode(ctx context.Context, code string) (*model.Team, error)
	CreateTeam(ctx context.Context, team model.Team) error
	GetTeams(ctx context.Context, league string) ([]*model.Team, error)
}

type IGameLogRepo interface {
	Create(ctx context.Context, gameLog model.GameLog) error
	FindByGameID(ctx context.Context, gameID uint) ([]*model.GameLog, error)
}

type IGameRepo interface {
	FindGamesByDate(ctx context.Context, date time.Time) ([]*model.Game, error)
	FindGameIDByDate(ctx context.Context, date time.Time, homeTeam, season string) (uint, error)
	Create(ctx context.Context, game *model.Game) (uint, error)
}

type Repositories struct {
	TeamRepo    ITeamRepo
	GameRepo    IGameRepo
	PlayerRepo  IPlayerRepo
	GameLogRepo IGameLogRepo
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		TeamRepo:    NewTeamRepo(db),
		GameRepo:    NewGameRepo(db),
		PlayerRepo:  NewPlayerRepo(db),
		GameLogRepo: NewGameLogRepo(db),
	}
}
