package service

import (
	"context"
	"fmt"
	"time"

	"github.com/new-pop-corn/internal/pkg/model"
	"github.com/new-pop-corn/internal/pkg/repo"
	"github.com/opentracing/opentracing-go/log"
)

type GameService interface {
	GetGamesByDate(ctx context.Context, date time.Time) ([]*model.GameResp, error)
	// GetGameByGame(ctx context.Context, gameID uint) (*model.Game, error)
	GetGameLogByGameID(ctx context.Context, gameID uint) (map[string]*model.PlayerActionLog, error)
}

type gameService struct {
	GameRepo   repo.GameRepo
	TeamRepo   repo.TeamRepo
	PlayerRepo repo.PlayerRepo
}

func NewGameService(gp repo.GameRepo, tp repo.TeamRepo, pp repo.PlayerRepo) GameService {
	return &gameService{
		GameRepo:   gp,
		TeamRepo:   tp,
		PlayerRepo: pp,
	}
}

func (gs *gameService) GetGamesByDate(ctx context.Context, date time.Time) ([]*model.GameResp, error) {
	var gamesResp []*model.GameResp
	games, err := gs.GameRepo.GetGamesByDate(ctx, date)
	if err != nil {
		return nil, err
	}
	for _, game := range games {
		homeTeam, err := gs.TeamRepo.GetTeamByCode(ctx, game.HomeTeam)
		if err != nil {
			return nil, err
		}
		awayTeam, err := gs.TeamRepo.GetTeamByCode(ctx, game.AwayTeam)
		if err != nil {
			return nil, err
		}
		gameResp := &model.GameResp{
			Game:         *game,
			HomeTeamIcon: homeTeam.TeamLogo,
			AwayTeamIcon: awayTeam.TeamLogo,
		}
		gamesResp = append(gamesResp, gameResp)
	}
	return gamesResp, nil
}

func (gs *gameService) GetGameLogByGameID(ctx context.Context, gameID uint) (map[string]*model.PlayerActionLog, error) {
	eventGroupByPlayer := map[string]*model.PlayerActionLog{}
	gameLogs, err := gs.GameRepo.GetGameLogByGameID(ctx, gameID)

	if err != nil {
		return nil, err
	}

	for _, gameLog := range gameLogs {
		player, err := gs.PlayerRepo.GetPlayerByID(ctx, gameLog.PlayerID)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		team, err := gs.TeamRepo.GetTeamByID(ctx, player.TeamID)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		if _, ok := eventGroupByPlayer[player.Name]; !ok {
			eventGroupByPlayer[player.Name] = &model.PlayerActionLog{
				Event: map[int][]model.Action{},
				Pts:   map[int]uint{},
				Team:  team.Name,
			}
		}
		action := &model.Action{
			TimeStamp: convertTo8601(gameLog.TS),
			Action:    gameLog.Event,
		}

		eventGroupByPlayer[player.Name].Event[gameLog.Quarter] = append(eventGroupByPlayer[player.Name].Event[gameLog.Quarter], *action)
		eventGroupByPlayer[player.Name].Pts[gameLog.Quarter] += cal(gameLog.Event)
	}

	return eventGroupByPlayer, nil
}

// func periodTime(ts float64) {
// 	min := 0.0
// 	sec := 0.0
// 	p := 0.0
// 	if ts <= 48 {
// 		p = math.Ceil(ts / 12)
// 		min = math.Floor((p * 12) - ts)
// 		sec = math.Round(((p * 12) - ts - min) * 60)
// 	} else {
// 		ts = ts - 48
// 		p = math.Ceil(ts / 5)
// 		min = math.Floor((p * 5) - ts)
// 		sec = math.Round(((p * 5) - ts - min) * 60)
// 	}
// 	return ('0' + min).slice(-2) + ':' + ('0' + sec).slice(-2)
// }

func convertTo8601(ts float32) string {
	minutes := int(ts)
	seconds := int((float64(ts) - float64(minutes)) * 60)
	timeString := fmt.Sprintf("%02d:%02d:00", minutes, seconds)
	return timeString
}

func cal(event model.Event) uint {
	switch event {
	case model.MAKE3:
		return 3
	case model.MAKE2:
		return 2
	case model.MAKEFT:
		return 1
	default:
		return 0
	}
}
