package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/new-pop-corn/config"
	"github.com/new-pop-corn/internal/api"
	"github.com/new-pop-corn/internal/app/initialize"
	db "github.com/new-pop-corn/internal/database"
	"github.com/new-pop-corn/internal/repo"
	"github.com/new-pop-corn/internal/service"
	"go.uber.org/zap"
)

func Run() {

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	//init repo
	repos := repo.NewRepositories(db.Get())
	//init service
	services := service.NewServices(service.Deps{
		Repos: repos,
	})
	//init router
	engine := gin.Default()
	api.NewHandler(&api.Config{
		R:               engine,
		Services:        services,
		TimeoutDuration: time.Duration(config.ServerConf.Server.HTTPTimeout) * time.Second,
	})

	srv := &http.Server{
		Addr:    config.ServerConf.Server.Addr,
		Handler: engine,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatalf("listen: %s\n", err)
		}
	}()
	zap.S().Infof("Server listen on: %s", config.ServerConf.Server.Addr)

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("Shutdown Server ...")

	timeout := config.ServerConf.Server.ShutdownTimeout

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of ${timeout} seconds.
	<-ctx.Done()
	zap.S().Infof("timeout of %d seconds.", timeout)
}
