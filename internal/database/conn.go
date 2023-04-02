package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/new-pop-corn/config"
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

	c := config.ServerConf.MySQLInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Name)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}
	migration.Migrate(db)
	return nil
}

func Get() *gorm.DB {
	return db
}
