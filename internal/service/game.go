package service

import (
	"context"
	"time"

	"github.com/new-pop-corn/internal/model"
	"github.com/new-pop-corn/internal/repo"
	"go.uber.org/zap"
)

type gameService struct {
	repo repo.Repositories
}

func NewGameService(repo repo.Repositories) GameService {
	return &gameService{
		repo: repo,
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
			zap.S().Error(err)
			return nil, err
		}
		team, err := gs.repo.TeamRepo.FindByID(ctx, player.TeamID)
		if err != nil {
			zap.S().Error(err)
			return nil, err
		}
		formatGameLog(game, gl, playerLogs, player, team)

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

	return result, nil
}

func formatGameLog(game *model.Game, gl *model.GameLog,
	playerLogs map[string]*model.PlayerActionLog, player *model.Player, team *model.Team) {
	if _, ok := playerLogs[player.Name]; !ok {
		playerLogs[player.Name] = &model.PlayerActionLog{
			Event:      map[int][]model.Action{},
			Team:       team.Name,
			Durations:  make([][]map[model.Event]float32, 5),
			Period:     gl.Period,
			IsHomeTeam: team.Code == game.HomeTeam,
			TotalTime:  0,
		}

		switch gl.Event {
		case model.OFF:
			playerLogs[player.Name].IsOnField = 0
			start := float32(720)
			end := gl.RemainingSecondsInPeriod
			playerLogs[player.Name].Durations[gl.Period] = []map[model.Event]float32{{model.ON: start},
				{model.OFF: end}}
		case model.ON:
			playerLogs[player.Name].IsOnField = 1
			start := gl.RemainingSecondsInPeriod
			playerLogs[player.Name].Durations[gl.Period] = []map[model.Event]float32{{model.ON: start}}
		default:
			playerLogs[player.Name].IsOnField = 1
		}
	} else {

		switch gl.Event {
		case model.OFF:
			if playerLogs[player.Name].IsOnField == 1 {
				playerLogs[player.Name].IsOnField = 0
				end := gl.RemainingSecondsInPeriod
				playerLogs[player.Name].Durations[gl.Period] = append(playerLogs[player.Name].Durations[gl.Period],
					map[model.Event]float32{model.OFF: end})
			}
		case model.ON:
			if playerLogs[player.Name].IsOnField == 0 {
				playerLogs[player.Name].IsOnField = 1
				start := gl.RemainingSecondsInPeriod
				playerLogs[player.Name].Durations[gl.Period] = append(playerLogs[player.Name].Durations[gl.Period],
					map[model.Event]float32{model.ON: start})
			} else if playerLogs[player.Name].IsOnField == 1 && playerLogs[player.Name].Period != gl.Period {
				start := gl.RemainingSecondsInPeriod
				playerLogs[player.Name].Durations[gl.Period] = append(playerLogs[player.Name].Durations[gl.Period],
					map[model.Event]float32{model.ON: start})
			}
		}
		playerLogs[player.Name].Period = gl.Period
	}

}
