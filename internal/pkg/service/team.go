package service

import (
	"context"

	"github.com/new-pop-corn/internal/pkg/model"
	"github.com/new-pop-corn/internal/pkg/repo"
)

type TeamService interface {
	GetTeamByName(ctx context.Context, teamName string) (*model.Team, error)
	GetTeamByID(ctx context.Context, teamID uint) (*model.Team, error)
	GetTeams(ctx context.Context, league string) ([]*model.Team, error)
}

type teamService struct {
	TeamRepository repo.TeamRepository
}

func NewTeamService(tr repo.TeamRepository) TeamService {
	return &teamService{
		TeamRepository: tr,
	}
}

func (ts *teamService) GetTeamByName(ctx context.Context, teamName string) (*model.Team, error) {
	product, err := ts.TeamRepository.GetTeamByName(ctx, teamName)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ts *teamService) GetTeamByID(ctx context.Context, teamID uint) (*model.Team, error) {
	return ts.TeamRepository.GetTeamByID(ctx, teamID)
}

func (ts *teamService) GetTeams(ctx context.Context, league string) ([]*model.Team, error) {
	return ts.TeamRepository.GetTeams(ctx, league)
}
