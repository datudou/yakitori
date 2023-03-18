package migration

import (
	"fmt"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/new-pop-corn/internal/pkg/model"
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
				ID: "20230317",
				Migrate: func(tx *gorm.DB) error {
					return initData(tx)
				},
			},
		})

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	fmt.Println("Database migration did run successfully")
}
