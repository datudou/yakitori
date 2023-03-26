package repo

import (
	"context"
	"time"

	"github.com/new-pop-corn/internal/model"
	"gorm.io/gorm"
)

type DBOption func(*gorm.DB) *gorm.DB

type gameRepo struct {
	DB *gorm.DB
}

// FindByDate implements IGameRepo

func NewGameRepo(db *gorm.DB) IGameRepo {
	return &gameRepo{
		DB: db,
	}
}

func (gr *gameRepo) Create(ctx context.Context, game *model.Game) (uint, error) {
	if err := gr.DB.Create(&game).Error; err != nil {
		return 0, err
	}
	return game.ID, nil
}

func (gp *gameRepo) FindGamesByDate(ctx context.Context, date time.Time) ([]*model.Game, error) {
	var games []*model.Game
	err := gp.DB.Table("games").Where("date = ?", date).Find(&games).Error
	if err != nil {
		return nil, err
	}
	return games, nil
}

func (gp *gameRepo) FindGameIDByDate(ctx context.Context, date time.Time, homeTeam, season string) (uint, error) {
	var game *model.Game
	err := gp.DB.Table("games").Select("id").
		Where("start_date = ? and home_team = ? and season = ? ", date, homeTeam, season).
		First(&game).Error
	if err != nil {
		return 0, err
	}
	return game.ID, nil
}
