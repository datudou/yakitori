package initialize

import (
	db "github.com/new-pop-corn/internal/database"
	"go.uber.org/zap"
)

func InitDB() {
	if err := db.SetupConn(); err != nil {
		zap.S().Fatal(err)
	}
}
