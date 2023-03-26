package migration

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/new-pop-corn/internal/model"
	"github.com/new-pop-corn/internal/repo"

	"github.com/new-pop-corn/internal/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type dataInitial struct {
	PR repo.IPlayerRepo
	TR repo.ITeamRepo
}

func NewDataInitial(db *gorm.DB) *dataInitial {
	return &dataInitial{
		PR: repo.NewPlayerRepo(db),
		TR: repo.NewTeamRepo(db),
	}
}

func (di *dataInitial) initPlayerData() error {
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
			team, err := di.TR.GetTeamByCode(context.Background(), strings.ToUpper(r.Team))
			if err != nil {
				return err
			}

			player := model.Player{
				SimpleName: simpleName(r.Name),
				Name:       r.Name,
				TeamID:     team.ID,
				Sport:      "NBA",
				Position:   r.Pos,
				Age:        r.Age,
			}
			_, err = di.PR.CreatePlayer(context.Background(), player)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (di *dataInitial) initTeamData() error {
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
			err = di.TR.CreateTeam(context.Background(), team)
			if err != nil {
				return err
			}

		}
	}
	return nil
}

func simpleName(name string) string {
	parts := strings.Split(name, " ")
	firstInitial := string(parts[0][0])
	abbreviated := fmt.Sprintf("%s. %s", firstInitial, parts[1])
	return abbreviated
}
