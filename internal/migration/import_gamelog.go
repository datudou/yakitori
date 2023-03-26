package migration

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/new-pop-corn/internal/model"
	"github.com/new-pop-corn/internal/repo"
	"github.com/new-pop-corn/internal/utils"
	"github.com/new-pop-corn/internal/utils/gamelog"
	"gorm.io/gorm"
)

func importGame(season, date string, db *gorm.DB) error {
	dirPath := fmt.Sprintf("./assets/nba-game-log/%s/%s", season, date)
	repos := repo.NewRepositories(db)
	ctx := context.Background()

	files, err := utils.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, f := range files {
		game, err := genGame(f.Name(), season)
		if err != nil {
			return err
		}
		var gameID uint
		gameID, err = repos.GameRepo.FindGameIDByDate(ctx, game.StartDate, game.HomeTeam, game.Season)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				gameID, _ = repos.GameRepo.Create(context.Background(), game)
			} else {
				return err
			}
		}

		var result []GameLogRecord
		readGameLogJson(fmt.Sprintf("%s/%s", dirPath, f.Name()), &result)

		for _, r := range result {
			playerActions := parseGameLog(r.Description)
			for _, pa := range playerActions {

				if pa.Player == "" || pa.Event == "" {
					continue
				}

				awayTeam, err := repos.TeamRepo.FindByCode(ctx, game.AwayTeam)
				if err != nil {
					fmt.Printf("AwayTeam %s not found", game.AwayTeam)
					return err
				}
				homeTeam, err := repos.TeamRepo.FindByCode(ctx, game.HomeTeam)
				if err != nil {
					fmt.Printf("HomeTeam %s not found", game.HomeTeam)
					return err
				}
				var player *model.Player
				player, err = repos.PlayerRepo.FindBySimpleNameAndTeamID(ctx, pa.Player, awayTeam.ID)

				if err != nil {
					if err == gorm.ErrRecordNotFound {
						player, err = repos.PlayerRepo.FindBySimpleNameAndTeamID(ctx, pa.Player, homeTeam.ID)
						if err != nil {
							fmt.Printf("Player %s not found", pa.Player)
							return err
						}
					} else {
						fmt.Printf("Player %s not found", pa.Player)
						return err
					}
				}

				uuid := uuid.NewMD5(uuid.Nil, []byte(fmt.Sprintf("%d-%d-%s-%f-%s-%d", gameID,
					player.ID, pa.Event, r.RemainingSecondsInPeriod, r.Description, r.HomeScore)))
				gameLog := model.GameLog{
					GameID:                   gameID,
					PlayerID:                 player.ID,
					Event:                    pa.Event,
					RemainingSecondsInPeriod: r.RemainingSecondsInPeriod,
					Period:                   r.Period,
					PeriodType:               "quarter",
					AwayScore:                r.AwayScore,
					HomeScore:                r.HomeScore,
					Description:              r.Description,
					EventID:                  uuid.String(),
				}

				err = repos.GameLogRepo.Create(context.Background(), gameLog)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GameLogRecord struct {
	Period                   int     `json:"period"`
	RemainingSecondsInPeriod float32 `json:"remaining_seconds_in_period"`
	Team                     string  `json:"team"`
	PeriodDuration           float32 `json:"period_duration"`
	Description              string  `json:"description"`
	AwayScore                int     `json:"away_score"`
	HomeScore                int     `json:"home_score"`
	RelevantTeam             string  `json:"relevant_team"`
}

func readGameLogJson(path string, result *[]GameLogRecord) error {

	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}

	byteValute, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	json.Unmarshal(byteValute, &result)
	return nil
}

func parseGameLog(desc string) []*model.PlayerAction {
	var pas []*model.PlayerAction
	splitDesc := strings.Split(desc, "(")
	for i := 0; i < len(splitDesc); i++ {
		pas = append(pas, gamelog.Parse(splitDesc[i])...)
	}
	return pas
}

func genGame(file, season string) (*model.Game, error) {
	fileName := strings.Split(file, ".")[0]
	fileNameWithoutDate := strings.Split(fileName, "|")[1]
	date := strings.Split(fileName, "|")[0]
	homeTeam := strings.Split(fileNameWithoutDate, "vs")[1]
	awayTeam := strings.Split(fileNameWithoutDate, "vs")[0]
	hss := strings.Split(homeTeam, "-")
	ass := strings.Split(awayTeam, "-")
	homeTeamScore, _ := strconv.ParseInt(hss[1], 10, 0)
	awyaTeamScore, _ := strconv.ParseInt(ass[1], 10, 0)
	layout := "2006-01-02 15:04:05+00:00"
	t, err := time.Parse(layout, date)
	if err != nil {
		return nil, err
	}

	game := model.Game{
		Season:        season,
		HomeTeam:      hss[0],
		AwayTeam:      ass[0],
		HomeTeamScore: int(homeTeamScore),
		AwayTeamScore: int(awyaTeamScore),
		StartDate:     t,
	}
	return &game, nil
}
