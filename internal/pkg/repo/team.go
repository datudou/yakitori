package repo

import (
	"context"

	"github.com/new-pop-corn/internal/pkg/model"
	"gorm.io/gorm"
)

type TeamRepository interface {
	GetTeamByID(ctx context.Context, teamID uint) (*model.Team, error)
	GetTeamByName(ctx context.Context, name string) (*model.Team, error)
	Create(ctx context.Context, t model.Team) (*model.Team, error)
	GetTeams(ctx context.Context, league string) ([]*model.Team, error)
}

type teamRepository struct {
	DB *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &teamRepository{
		DB: db,
	}
}

func (tr *teamRepository) Create(ctx context.Context, team model.Team) (*model.Team, error) {
	if err := tr.DB.Create(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (tr *teamRepository) GetTeamByID(ctx context.Context, teamID uint) (*model.Team, error) {
	return getTeamByID(tr.DB, teamID)
}

func (tr *teamRepository) GetTeamByName(ctx context.Context, name string) (*model.Team, error) {
	var team model.Team
	err := tr.DB.Table("teams").Where("name=?", name).First(&team).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (tr *teamRepository) GetTeams(ctx context.Context, league string) ([]*model.Team, error) {
	var teams []*model.Team
	err := tr.DB.Table("teams").Where("league = ?", league).Find(&teams).Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}

// private function
func getTeamByID(db *gorm.DB, teamID uint) (*model.Team, error) {
	var team model.Team
	err := db.Table("teams").Select("name").Where("id=?", teamID).First(&team).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}
