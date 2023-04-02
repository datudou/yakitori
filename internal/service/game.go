package service

import (
	"context"
	"time"

	"github.com/golang/groupcache/lru"
	"github.com/new-pop-corn/internal/model"
	"github.com/new-pop-corn/internal/repo"
	"github.com/opentracing/opentracing-go/log"
)

type gameService struct {
	repo     repo.Repositories
	lruCache *lru.Cache
}

func NewGameService(repo repo.Repositories) GameService {
	return &gameService{
		repo:     repo,
		lruCache: lru.New(100),
	}
}

func (gs *gameService) GetGamesByDate(ctx context.Context, date time.Time) ([]*model.GameResp, error) {
	var gamesResp []*model.GameResp
	games, err := gs.repo.GameRepo.FindGamesByDate(ctx, date)
	if err != nil {
		return nil, err
	}
	for _, game := range games {
		homeTeam, err := gs.repo.TeamRepo.FindByCode(ctx, game.HomeTeam)
		if err != nil {
			return nil, err
		}
		awayTeam, err := gs.repo.TeamRepo.FindByCode(ctx, game.AwayTeam)
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

func (gs *gameService) GetGameLogByGameID(ctx context.Context, gameID uint) ([]*model.PlayerActionLog, error) {
	if v, ok := gs.lruCache.Get(gameID); ok {
		return v.([]*model.PlayerActionLog), nil
	}
	gls, err := gs.repo.GameLogRepo.Find(ctx, gs.repo.GameLogRepo.WithByGameID(gameID))
	if err != nil {
		return nil, err
	}

	game, err := gs.repo.GameRepo.FindByID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	playerLogs := make(map[string]*model.PlayerActionLog)
	for _, gl := range gls {
		player, err := gs.repo.PlayerRepo.FindByID(ctx, gl.PlayerID)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		team, err := gs.repo.TeamRepo.FindByID(ctx, player.TeamID)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		if _, ok := playerLogs[player.Name]; !ok {
			playerLogs[player.Name] = &model.PlayerActionLog{
				Event:      map[int][]model.Action{},
				Team:       team.Name,
				Durations:  make([][]map[model.Event]float32, 5),
				Period:     gl.Period,
				IsHomeTeam: team.Code == game.HomeTeam,
				TotalTime:  0,
			}
			if gl.Event == model.OFF {
				playerLogs[player.Name].IsOnField = 0
				playerLogs[player.Name].Durations[gl.Period] = []map[model.Event]float32{{model.ON: 720},
					{model.OFF: gl.RemainingSecondsInPeriod}}

			} else if gl.Event == model.ON {
				playerLogs[player.Name].IsOnField = 1
				playerLogs[player.Name].Durations[gl.Period] = []map[model.Event]float32{{model.ON: gl.RemainingSecondsInPeriod}}
			}

			if gl.Event != model.OFF {
				playerLogs[player.Name].IsOnField = 1
			}
		} else {
			if gl.Event == model.OFF && playerLogs[player.Name].IsOnField == 1 {
				playerLogs[player.Name].IsOnField = 0
				playerLogs[player.Name].Durations[gl.Period] = append(playerLogs[player.Name].Durations[gl.Period],
					map[model.Event]float32{model.OFF: gl.RemainingSecondsInPeriod})
			} else if gl.Event == model.ON && playerLogs[player.Name].IsOnField == 0 {
				playerLogs[player.Name].IsOnField = 1
				playerLogs[player.Name].Durations[gl.Period] = append(playerLogs[player.Name].Durations[gl.Period],
					map[model.Event]float32{model.ON: gl.RemainingSecondsInPeriod})
			} else if gl.Event == model.ON && playerLogs[player.Name].IsOnField == 1 && playerLogs[player.Name].Period != gl.Period {
				playerLogs[player.Name].Durations[gl.Period] = append(playerLogs[player.Name].Durations[gl.Period],
					map[model.Event]float32{model.ON: gl.RemainingSecondsInPeriod})
			}
			playerLogs[player.Name].Period = gl.Period

		}

		action := &model.Action{
			TimeStamp: gl.RemainingSecondsInPeriod,
			Action:    gl.Event,
		}

		playerLogs[player.Name].Event[gl.Period] = append(playerLogs[player.Name].Event[gl.Period], *action)
	}

	result := make([]*model.PlayerActionLog, 0, len(playerLogs))
	for k, v := range playerLogs {
		v.PlayerName = k
		result = append(result, v)
	}

	gs.lruCache.Add(gameID, result)
	return result, nil
}
