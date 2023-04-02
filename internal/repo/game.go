package repo

import (
	"context"
	"time"

	"github.com/new-pop-corn/internal/model"
	"gorm.io/gorm"
)

type gameRepo struct {
	DB *gorm.DB
}

func NewGameRepo(db *gorm.DB) IGameRepo {
	return &gameRepo{
		DB: db,
	}
}
func (gr *gameRepo) FindByID(ctx context.Context, gameID uint) (*model.Game, error) {
	var game *model.Game
	err := gr.DB.Table("games").
		Where("id = ?", gameID).
		First(&game).Error
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (gr *gameRepo) Create(ctx context.Context, game *model.Game) (uint, error) {
	if err := gr.DB.Create(&game).Error; err != nil {
		return 0, err
	}
	return game.ID, nil
}

func (gp *gameRepo) FindGamesByDate(ctx context.Context, date time.Time) ([]*model.Game, error) {
	var games []*model.Game
	err := gp.DB.Table("games").Where("start_date between ? and ?", date, date.Add(24*time.Hour)).Find(&games).Error
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
