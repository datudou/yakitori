package repo

import (
	"context"

	"github.com/new-pop-corn/internal/pkg/model"
	"gorm.io/gorm"
)

type PlayerRepo interface {
	GetPlayersByTeamID(ctx context.Context, teamID uint) (*model.Player, error)
	GetPlayerByName(ctx context.Context, playerName string) (*model.Player, error)
	GetPlayers(ctx context.Context, league string) ([]*model.Player, error)
	CreatePlayer(ctx context.Context, player model.Player) (*model.Player, error)
}

type playerRepo struct {
	DB *gorm.DB
}

func NewPlayerRepo(db *gorm.DB) PlayerRepo {
	return &playerRepo{
		DB: db,
	}
}

func (pr *playerRepo) GetPlayerByName(ctx context.Context, playerName string) (*model.Player, error) {
	var player model.Player
	err := pr.DB.Table("players").Where("name = ?", playerName).First(&player).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (pr *playerRepo) GetPlayersByTeamID(ctx context.Context, teamID uint) (*model.Player, error) {
	var player model.Player
	err := pr.DB.Table("players").Where("team_id = ?", teamID).First(&player).Error
	if err != nil {
		return nil, err
	}
	return &player, nil
}

func (pr *playerRepo) GetPlayers(ctx context.Context, league string) ([]*model.Player, error) {
	var players []*model.Player
	err := pr.DB.Table("players").Where("league = ?", league).Find(&players).Error
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (pr *playerRepo) CreatePlayer(ctx context.Context, player model.Player) (*model.Player, error) {
	if err := pr.DB.Create(&player).Error; err != nil {
		return nil, err
	}
	return &player, nil
}
