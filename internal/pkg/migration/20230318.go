package migration

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/new-pop-corn/internal/pkg/model"
	"github.com/new-pop-corn/internal/pkg/repo"

	"github.com/new-pop-corn/internal/pkg/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func initData(db *gorm.DB) error {
	dirPath := "./assets/team-meta-data"
	tr := repo.NewTeamRepository(db)
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
			_, err := tr.Create(context.Background(), team)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
