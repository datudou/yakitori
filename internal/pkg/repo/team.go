package repo

import (
	"context"

	"github.com/new-pop-corn/internal/pkg/model"
	"gorm.io/gorm"
)

type TeamRepo interface {
	GetTeamByID(ctx context.Context, teamID uint) (*model.Team, error)
	GetTeamByName(ctx context.Context, name string) (*model.Team, error)
	GetTeamByCode(ctx context.Context, code string) (*model.Team, error)
	// CreateTeams(ctx context.Context, teams []*model.Team) error
	CreateTeam(ctx context.Context, team model.Team) error
	GetTeams(ctx context.Context, league string) ([]*model.Team, error)
}

type teamRepo struct {
	DB *gorm.DB
}

func NewTeamRepo(db *gorm.DB) TeamRepo {
	return &teamRepo{
		DB: db,
	}
}

// func (tr *teamRepo) CreateTeams(ctx context.Context, teams []*model.Team) error {
// 	if err := tr.DB.CreateInBatches(&teams, 100).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func (tr *teamRepo) CreateTeam(ctx context.Context, team model.Team) error {
	if err := tr.DB.Create(&team).Error; err != nil {
		return err
	}
	return nil
}

func (tr *teamRepo) GetTeamByID(ctx context.Context, teamID uint) (*model.Team, error) {
	return getTeamByID(tr.DB, teamID)
}

func (tr *teamRepo) GetTeamByName(ctx context.Context, name string) (*model.Team, error) {
	var team model.Team
	err := tr.DB.Table("teams").Where("name = ?", name).First(&team).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (tr *teamRepo) GetTeamByCode(ctx context.Context, code string) (*model.Team, error) {
	var team model.Team
	err := tr.DB.Table("teams").Where("code = ?", code).First(&team).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (tr *teamRepo) GetTeams(ctx context.Context, league string) ([]*model.Team, error) {
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
	err := db.Table("teams").Where("id=?", teamID).First(&team).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}