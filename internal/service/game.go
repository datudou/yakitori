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

func (gs *gameService) GetGameLogByGameID(ctx context.Context, gameID uint) (map[string]*model.PlayerActionLog, error) {
	if v, ok := gs.lruCache.Get(gameID); ok {
		return v.(map[string]*model.PlayerActionLog), nil
	}
	egp := map[string]*model.PlayerActionLog{}
	gls, err := gs.repo.GameLogRepo.Find(ctx, gs.repo.GameLogRepo.WithByGameID(gameID))

	if err != nil {
		return nil, err
	}

	for _, gl := range gls {
		p, err := gs.repo.PlayerRepo.FindByID(ctx, gl.PlayerID)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		team, err := gs.repo.TeamRepo.FindByID(ctx, p.TeamID)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		if _, ok := egp[p.Name]; !ok {
			egp[p.Name] = &model.PlayerActionLog{
				Event:  map[int][]model.Action{},
				Team:   team.Name,
				Mins:   map[int][]map[model.Event]float32{},
				Period: gl.Period,
			}
			if gl.Event == model.OFF {
				egp[p.Name].IsOnField = 0
				egp[p.Name].Mins[gl.Period] = []map[model.Event]float32{{model.ON: 719}, {model.OFF: gl.RemainingSecondsInPeriod}}
			} else if gl.Event == model.ON {
				egp[p.Name].IsOnField = 1
				egp[p.Name].Mins[gl.Period] = []map[model.Event]float32{{model.ON: gl.RemainingSecondsInPeriod}}
			}

			if gl.Event != model.OFF {
				egp[p.Name].IsOnField = 1
			}

		} else {
			if gl.Event == model.OFF && egp[p.Name].IsOnField == 1 {
				egp[p.Name].IsOnField = 0
				egp[p.Name].Mins[gl.Period] = append(egp[p.Name].Mins[gl.Period],
					map[model.Event]float32{model.OFF: gl.RemainingSecondsInPeriod})
			} else if gl.Event == model.ON && egp[p.Name].IsOnField == 0 {
				egp[p.Name].IsOnField = 1
				egp[p.Name].Mins[gl.Period] = append(egp[p.Name].Mins[gl.Period],
					map[model.Event]float32{model.ON: gl.RemainingSecondsInPeriod})
			} else if gl.Event == model.ON && egp[p.Name].IsOnField == 1 && egp[p.Name].Period != gl.Period {
				egp[p.Name].Mins[gl.Period] = append(egp[p.Name].Mins[gl.Period],
					map[model.Event]float32{model.ON: gl.RemainingSecondsInPeriod})
			}
			egp[p.Name].Period = gl.Period
		}

		action := &model.Action{
			TimeStamp: gl.RemainingSecondsInPeriod,
			Action:    gl.Event,
		}

		egp[p.Name].Event[gl.Period] = append(egp[p.Name].Event[gl.Period], *action)
	}
	gs.lruCache.Add(gameID, egp)
	return egp, nil
}
