package service

import (
	"context"

	"github.com/new-pop-corn/internal/model"
	"github.com/new-pop-corn/internal/repo"
)

type teamService struct {
	repo repo.ITeamRepo
}

func NewTeamService(tr repo.ITeamRepo) TeamService {
	return &teamService{
		repo: tr,
	}
}

func (ts *teamService) GetTeamByName(ctx context.Context, teamName string) (*model.Team, error) {
	return ts.repo.FindByName(ctx, teamName)
}

func (ts *teamService) GetTeamByID(ctx context.Context, teamID uint) (*model.Team, error) {
	return ts.repo.FindByID(ctx, teamID)
}

func (ts *teamService) GetTeams(ctx context.Context, league string) ([]*model.Team, error) {
	return ts.repo.FindByLeague(ctx, league)
}
