package service

import (
	"context"
	"time"

	"github.com/new-pop-corn/internal/pkg/model"
	"github.com/new-pop-corn/internal/pkg/repo"
)

type GameService interface {
	GetGamesByDate(ctx context.Context, date time.Time) ([]*model.Game, error)
	// GetGameByGame(ctx context.Context, gameID uint) (*model.Game, error)
	GetGameLogByGameID(ctx context.Context, gameID uint) ([]*model.GameLog, error)
}

type gameService struct {
	GameRepo repo.GameRepo
}

func NewGameService(gp repo.GameRepo) GameService {
	return &gameService{
		GameRepo: gp,
	}
}

func (gs *gameService) GetGamesByDate(ctx context.Context, date time.Time) ([]*model.Game, error) {
	return gs.GameRepo.GetGamesByDate(ctx, date)
}

func (gs *gameService) GetGameLogByGameID(ctx context.Context, gameID uint) ([]*model.GameLog, error) {
	return gs.GameRepo.GetGameLogByGameID(ctx, gameID)
}
