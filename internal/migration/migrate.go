package migration

import (
	"fmt"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/new-pop-corn/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	m := gormigrate.New(
		db,
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			{
				ID: "Initial",
				Migrate: func(tx *gorm.DB) error {
					return tx.AutoMigrate(&model.Game{}, &model.Team{}, &model.Player{}, &model.Game{}, &model.GameLog{})
				},
				Rollback: func(tx *gorm.DB) error {
					return tx.Migrator().DropTable("game", "team", "player", "game_log")
				},
			},
			{
				ID: "inital_data",
				Migrate: func(tx *gorm.DB) error {
					di := NewDataInitial(tx)
					err := di.initTeamData()
					if err != nil {
						return err
					}
					err = di.initPlayerData()
					if err != nil {
						fmt.Printf("Error: %v", err)
						return err
					}
					return nil
				},
			},
			{
				ID: "gamelog-2022-10-18",
				Migrate: func(tx *gorm.DB) error {
					err := importGame("2022-2023", "2022-10-18", tx)
					if err != nil {
						return err
					}
					fmt.Println("finished import 2022-10-18")
					return nil
				},
			},
			{
				ID: "gamelog-2022-10-19",
				Migrate: func(tx *gorm.DB) error {
					return importGame("2022-2023", "2022-10-19", tx)
				},
			},
		})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	fmt.Println("Database migration did run successfully")
}
