package service

import (
	"context"
	"time"

	"github.com/new-pop-corn/internal/model"
	"github.com/new-pop-corn/internal/repo"
)

type GameService interface {
	GetGamesByDate(ctx context.Context, date time.Time) ([]*model.GameResp, error)
	// GetGameByGame(ctx context.Context, gameID uint) (*model.Game, error)
	GetGameLogByGameID(ctx context.Context, gameID uint) (map[string]*model.PlayerActionLog, error)
}

type PlayerService interface {
	GetPlayersByTeamID(ctx context.Context, teamID uint) (*model.Player, error)
	// GetPlayerByGame(ctx context.Context, gameID uint) (*model.Player, error)
	GetPlayers(ctx context.Context, league string) ([]*model.Player, error)
}

type TeamService interface {
	GetTeamByName(ctx context.Context, teamName string) (*model.Team, error)
	GetTeamByID(ctx context.Context, teamID uint) (*model.Team, error)
	GetTeams(ctx context.Context, league string) ([]*model.Team, error)
}

type Services struct {
	GameService   GameService
	TeamService   TeamService
	PlayerService PlayerService
}

type Deps struct {
	Repos    *repo.Repositories
	Services *Services
}

func NewServices(deps Deps) *Services {
	ts := NewTeamService(deps.Repos.TeamRepo)
	ps := NewPlayerService(deps.Repos.PlayerRepo)
	gs := NewGameService(*deps.Repos)

	return &Services{
		GameService:   gs,
		PlayerService: ps,
		TeamService:   ts,
	}
}
