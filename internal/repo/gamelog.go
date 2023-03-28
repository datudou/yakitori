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

func (glr *gameLogRepo) Create(ctx context.Context, gameLog model.GameLog) error {
	if err := glr.DB.Create(&gameLog).Error; err != nil {
		return err
	}
	return nil
}

func (glr *gameLogRepo) Find(ctx context.Context, opts ...DBOption) ([]*model.GameLog, error) {

	db := glr.optionDB(ctx, opts...)
	var gamelogs []*model.GameLog
	err := db.Find(&gamelogs).Error
	if err != nil {
		return nil, err
	}
	return gamelogs, nil
}

func (glr *gameLogRepo) WithByGameID(gameID uint) DBOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table("game_logs").Where("game_id = ?", gameID)
	}
}

func (glr *gameLogRepo) optionDB(ctx context.Context, opts ...DBOption) *gorm.DB {
	db := glr.DB.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	return db
}
