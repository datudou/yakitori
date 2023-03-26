package repo

import (
	"context"

	"github.com/new-pop-corn/internal/model"
	"gorm.io/gorm"
)

type teamRepo struct {
	DB *gorm.DB
}

func NewTeamRepo(db *gorm.DB) ITeamRepo {
	return &teamRepo{
		DB: db,
	}
}

func (tr *teamRepo) Create(ctx context.Context, team model.Team) error {
	if err := tr.DB.Create(&team).Error; err != nil {
		return err
	}
	return nil
}

func (tr *teamRepo) FindByID(ctx context.Context, teamID uint) (*model.Team, error) {
	return getTeamByID(tr.DB, teamID)
}

func (tr *teamRepo) FindByName(ctx context.Context, name string) (*model.Team, error) {
	var team model.Team
	err := tr.DB.Table("teams").Where("name = ?", name).First(&team).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (tr *teamRepo) FindByCode(ctx context.Context, code string) (*model.Team, error) {
	var team model.Team
	err := tr.DB.Table("teams").Where("code = ?", code).First(&team).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func (tr *teamRepo) FindByLeague(ctx context.Context, league string) ([]*model.Team, error) {
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
