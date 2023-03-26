package repo

import (
	"context"

	"github.com/new-pop-corn/internal/model"
	"gorm.io/gorm"
)

type playerRepo struct {
	DB *gorm.DB
}

func NewPlayerRepo(db *gorm.DB) IPlayerRepo {
	return &playerRepo{
		DB: db,
	}
}

func (pr *playerRepo) FindByPlayerName(ctx context.Context, playerName string) (*model.Player, error) {
	var player model.Player
	err := pr.DB.Table("players").Where("name = ?", playerName).First(&player).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (pr *playerRepo) FindBySimpleNameAndTeamID(ctx context.Context, playerName string, teamID uint) (*model.Player, error) {
	var player model.Player
	err := pr.DB.Table("players").Where("simple_name = ? and team_id = ?", playerName, teamID).First(&player).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (pr *playerRepo) FindByID(ctx context.Context, id uint) (*model.Player, error) {
	var player model.Player
	err := pr.DB.Table("players").Where("id = ?", id).First(&player).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (pr *playerRepo) FindByTeamID(ctx context.Context, teamID uint) (*model.Player, error) {
	var player model.Player
	err := pr.DB.Table("players").Where("team_id = ?", teamID).First(&player).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (pr *playerRepo) FindByLeague(ctx context.Context, league string) ([]*model.Player, error) {
	var players []*model.Player
	err := pr.DB.Table("players").Where("league = ?", league).Find(&players).Error
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (pr *playerRepo) Create(ctx context.Context, player model.Player) (*model.Player, error) {
	if err := pr.DB.Create(&player).Error; err != nil {
		return nil, err
	}
	return &player, nil
}
