package model

import (
	"log"
	"os"
	"time"

	"github.com/new-pop-corn/internal/pkg/migration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB represents the database connection using gorm
var db *gorm.DB

func SetupConn() error {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             500 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Error,           // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                   // Disable color
		},
	)

	dsn := "host=localhost user=kenyiwang dbname=popcorn port=5432 sslmode=disable TimeZone=UTC"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}

	return nil
}

func DB() *gorm.DB {
	return db
}

func init() {
	err := SetupConn()
	if err != nil {
		log.Fatal(err)
	}
	migration.Migrate(db)
}
