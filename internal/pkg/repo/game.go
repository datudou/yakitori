package repo

import (
	"context"
	"time"

	"github.com/new-pop-corn/internal/pkg/model"
	"gorm.io/gorm"
)

type GameRepo interface {
	GetGamesByDate(ctx context.Context, date time.Time) ([]*model.Game, error)
	GetGameLogByGameID(ctx context.Context, gameID uint) ([]*model.GameLog, error)
	CreateGame(ctx context.Context, game model.Game) (uint, error)
	CreateGameLog(ctx context.Context, gameLog model.GameLog) error
}

type gameRepo struct {
	DB *gorm.DB
}

func NewGameRepo(db *gorm.DB) GameRepo {
	return &gameRepo{
		DB: db,
	}
}

// CreateGameLog implements GameRepo
func (gr *gameRepo) CreateGameLog(ctx context.Context, gameLog model.GameLog) error {
	if err := gr.DB.Create(&gameLog).Error; err != nil {
		return err
	}
	return nil
}

func (gr *gameRepo) GetGameLogByGameID(ctx context.Context, gameID uint) ([]*model.GameLog, error) {
	var gamelogs []*model.GameLog
	err := gr.DB.Table("game_logs").Where("game_id = ?", gameID).Find(&gamelogs).Error
	if err != nil {
		return nil, err
	}
	return gamelogs, nil
}

func (gr *gameRepo) CreateGame(ctx context.Context, game model.Game) (uint, error) {
	if err := gr.DB.Create(&game).Error; err != nil {
		return 0, err
	}
	return game.ID, nil
}

func (gp *gameRepo) GetGamesByDate(ctx context.Context, date time.Time) ([]*model.Game, error) {
	var games []*model.Game
	err := gp.DB.Table("games").Where("date = ?", date).Find(&games).Error
	if err != nil {
		return nil, err
	}
	return games, nil
}
