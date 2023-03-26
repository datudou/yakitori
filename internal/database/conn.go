package model

import (
	"log"
	"os"
	"time"

	"github.com/new-pop-corn/internal/migration"
	"gorm.io/driver/mysql"
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

	dsn := "root:@tcp(127.0.0.1:3306)/popcorn?charset=utf8mb4&parseTime=True"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
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
