package service

import (
	"context"

	"github.com/new-pop-corn/internal/model"
	"github.com/new-pop-corn/internal/repo"
)

type playerService struct {
	repo repo.IPlayerRepo
}

func NewPlayerService(pr repo.IPlayerRepo) PlayerService {
	return &playerService{
		repo: pr,
	}
}

func (ps *playerService) GetPlayersByTeamID(ctx context.Context, teamID uint) (*model.Player, error) {
	return ps.repo.FindByTeamID(ctx, teamID)
}

func (ps *playerService) GetPlayers(ctx context.Context, league string) ([]*model.Player, error) {
	return ps.repo.FindByLeague(ctx, league)
}
