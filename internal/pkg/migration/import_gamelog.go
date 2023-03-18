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

	"github.com/new-pop-corn/internal/pkg/model"
	"github.com/new-pop-corn/internal/pkg/repo"
	"github.com/new-pop-corn/internal/pkg/utils"
	"gorm.io/gorm"
)

func importGame(season, date string, db *gorm.DB) error {
	dirPath := fmt.Sprintf("./assets/nba-game-log/%s/%s", season, date)
	gr := repo.NewGameRepo(db)
	pr := repo.NewPlayerRepo(db)

	files, err := utils.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, f := range files {
		fileName := strings.Split(f.Name(), ".")[0]
		homeTeam := strings.Split(fileName, "vs")[1]
		awayTeam := strings.Split(fileName, "vs")[0]
		hss := strings.Split(homeTeam, "-")
		ass := strings.Split(awayTeam, "-")
		homeTeamScore, _ := strconv.ParseInt(hss[1], 10, 0)
		awyaTeamScore, _ := strconv.ParseInt(ass[1], 10, 0)
		layout := "2006-01-02"
		t, err := time.Parse(layout, date)
		if err != nil {
			return err
		}

		game := model.Game{
			Season:        season,
			HomeTeam:      hss[0],
			AwayTeam:      ass[0],
			HomeTeamScore: int(homeTeamScore),
			AwayTeamScore: int(awyaTeamScore),
			Date:          t,
		}

		gameID, _ := gr.CreateGame(context.Background(), game)

		jsonFile, err := os.Open(fmt.Sprintf("%s/%s", dirPath, f.Name()))
		if err != nil {
			return err
		}

		byteValute, err := io.ReadAll(jsonFile)
		if err != nil {
			return err
		}
		jsonFile.Close()
		type jsonRead struct {
			Quater    int         `json:"period"`
			TimeStamp float32     `json:"ts"`
			Team      string      `json:"team"`
			Player    string      `json:"player"`
			Event     model.Event `json:"event"`
		}
		var result []jsonRead

		json.Unmarshal(byteValute, &result)
		for _, r := range result {
			player, err := pr.GetPlayerByName(context.Background(), r.Player)
			if err != nil {
				return err
			}

			gameLog := model.GameLog{
				GameID:   gameID,
				PlayerID: player.ID,
				Event:    r.Event,
				TS:       r.TimeStamp,
				Quarter:  r.Quater,
			}
			err = gr.CreateGameLog(context.Background(), gameLog)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
