package repo

import (
	"context"

	"github.com/new-pop-corn/internal/model"
	"gorm.io/gorm"
)

type gameLogRepo struct {
	DB *gorm.DB
}

func NewGameLogRepo(db *gorm.DB) IGameLogRepo {
	return &gameLogRepo{
		DB: db,
	}

}

// CreateGameLog implements GameRepo
func (gr *gameLogRepo) Create(ctx context.Context, gameLog model.GameLog) error {
	if err := gr.DB.Create(&gameLog).Error; err != nil {
		return err
	}
	return nil
}

func (gr *gameLogRepo) FindByGameID(ctx context.Context, gameID uint) ([]*model.GameLog, error) {
	var gamelogs []*model.GameLog
	err := gr.DB.Table("game_logs").Where("game_id = ?", gameID).Find(&gamelogs).Error
	if err != nil {
		return nil, err
	}
	return gamelogs, nil
}
