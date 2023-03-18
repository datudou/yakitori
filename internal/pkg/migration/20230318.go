package migration

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/new-pop-corn/internal/pkg/model"
	"github.com/new-pop-corn/internal/pkg/repo"

	"github.com/new-pop-corn/internal/pkg/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type initData struct {
	PR repo.PlayerRepo
	TR repo.TeamRepo
}

func NewInitalData(db *gorm.DB) *initData {
	return &initData{
		PR: repo.NewPlayerRepo(db),
		TR: repo.NewTeamRepo(db),
	}
}

func (id *initData) initPlayerData() error {
	dirPath := "./assets/nba-player-data"
	files, err := utils.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
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
			Rank int     `json:"RANK"`
			Name string  `json:"NAME"`
			Team string  `json:"TEAM"`
			Pos  string  `json:"POS"`
			Age  float32 `json:"AGE"`
		}
		var result []jsonRead
		json.Unmarshal(byteValute, &result)
		for _, r := range result {
			team, err := id.TR.GetTeamByCode(context.Background(), strings.ToUpper(r.Team))
			if err != nil {
				return err
			}

			player := model.Player{
				Name:     r.Name,
				TeamID:   team.ID,
				Sport:    "NBA",
				Position: r.Pos,
				Age:      r.Age,
			}
			fmt.Printf("player: %v", player)
			_, err = id.PR.CreatePlayer(context.Background(), player)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (id *initData) initTeamData() error {
	dirPath := "./assets/team-meta-data"
	files, err := utils.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		jsonFile, err := os.Open(fmt.Sprintf("%s/%s", dirPath, f.Name()))
		if err != nil {
			return err
		}

		byteValute, err := io.ReadAll(jsonFile)
		if err != nil {
			return err
		}
		jsonFile.Close()

		var teams []model.Team
		json.Unmarshal(byteValute, &teams)
		for _, team := range teams {
			err = id.TR.CreateTeam(context.Background(), team)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
