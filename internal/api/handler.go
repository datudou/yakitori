package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/new-pop-corn/internal/middleware"
	"github.com/new-pop-corn/internal/model/apperrors"
	"github.com/new-pop-corn/internal/service"
	"github.com/sirupsen/logrus"
)

// Handler struct holds required services for handler to function
type Handler struct {
	Services *service.Services
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R               *gin.Engine
	Services        *service.Services
	TimeoutDuration time.Duration
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	h := &Handler{
		Services: c.Services,
	}

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	c.R.Use(gin.Recovery())
	c.R.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))

	g := c.R.Group("/api/v1")
	{
		g.GET("/teams/:league", h.getTeams)
		g.GET("/team/:id", h.getTeamByID)
		g.GET("/team/:id/players", h.getPlayersByTeamID)
		g.GET("/game/:id/gamelog", h.getGameLogByGameID)
		g.GET("/games/:date", h.getGamesByDate)
	}
}
