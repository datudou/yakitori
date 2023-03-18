package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/new-pop-corn/internal/pkg/api"
	db "github.com/new-pop-corn/internal/pkg/database"
	"github.com/new-pop-corn/internal/pkg/repo"
	"github.com/new-pop-corn/internal/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {

	// init db
	if err := db.SetupConn(); err != nil {
		log.Fatal(err)
	}
	engine := gin.Default()

	//init repo
	tr := repo.NewTeamRepo(db.DB())
	// pr := repo.NewPlayerRepo(db.DB())
	//init service
	ts := service.NewTeamService(tr)

	api.NewHandler(&api.Config{
		R:               engine,
		TeamService:     ts,
		TimeoutDuration: time.Duration(5 * time.Second),
	})
	logrus.SetLevel(logrus.DebugLevel)

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
