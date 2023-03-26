package repo

import (
	"context"
	"time"

	"github.com/new-pop-corn/internal/model"
	"gorm.io/gorm"
)

type IPlayerRepo interface {
	FindByTeamID(ctx context.Context, teamID uint) (*model.Player, error)
	FindByPlayerName(ctx context.Context, playerName string) (*model.Player, error)
	FindByID(ctx context.Context, id uint) (*model.Player, error)
	FindByLeague(ctx context.Context, league string) ([]*model.Player, error)
	FindBySimpleNameAndTeamID(ctx context.Context, simpleName string, teamID uint) (*model.Player, error)
	Create(ctx context.Context, player model.Player) (*model.Player, error)
}
type ITeamRepo interface {
	FindByID(ctx context.Context, teamID uint) (*model.Team, error)
	FindByName(ctx context.Context, name string) (*model.Team, error)
	FindByCode(ctx context.Context, code string) (*model.Team, error)
	FindByLeague(ctx context.Context, league string) ([]*model.Team, error)
	Create(ctx context.Context, team model.Team) error
}

type IGameLogRepo interface {
	FindByGameID(ctx context.Context, gameID uint) ([]*model.GameLog, error)
	Create(ctx context.Context, gameLog model.GameLog) error
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
