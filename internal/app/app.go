package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/new-pop-corn/internal/api"
	db "github.com/new-pop-corn/internal/database"
	"github.com/new-pop-corn/internal/repo"
	"github.com/new-pop-corn/internal/service"
	"github.com/sirupsen/logrus"
)

func Run() {
	// init db
	if err := db.SetupConn(); err != nil {
		logrus.Fatal(err)
	}
	logrus.SetLevel(logrus.DebugLevel)

	engine := gin.Default()
	//init repo
	repos := repo.NewRepositories(db.DB())

	//init service
	services := service.NewServices(service.Deps{
		Repos: repos,
	})

	//init api
	api.NewHandler(&api.Config{
		R:               engine,
		Services:        services,
		TimeoutDuration: time.Duration(5 * time.Second),
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")

}
