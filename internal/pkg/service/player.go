package service

import (
	"context"

	"github.com/new-pop-corn/internal/pkg/model"
	"github.com/new-pop-corn/internal/pkg/repo"
)

type PlayerService interface {
	GetPlayersByTeamID(ctx context.Context, teamID uint) (*model.Player, error)
	// GetPlayerByGame(ctx context.Context, gameID uint) (*model.Player, error)
	GetPlayers(ctx context.Context, league string) ([]*model.Player, error)
}

type playerService struct {
	PlayerRepo repo.PlayerRepo
}

func NewPlayerService(pr repo.PlayerRepo) PlayerService {
	return &playerService{
		PlayerRepo: pr,
	}
}

func (ps *playerService) GetPlayersByTeamID(ctx context.Context, teamID uint) (*model.Player, error) {
	return ps.PlayerRepo.GetPlayersByTeamID(ctx, teamID)
}

func (ps *playerService) GetPlayers(ctx context.Context, league string) ([]*model.Player, error) {
	return ps.PlayerRepo.GetPlayers(ctx, league)
}
