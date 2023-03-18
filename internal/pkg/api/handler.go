package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/new-pop-corn/internal/pkg/middleware"
	"github.com/new-pop-corn/internal/pkg/model/apperrors"
	"github.com/new-pop-corn/internal/pkg/service"
	"github.com/sirupsen/logrus"
)

// Handler struct holds required services for handler to function
type Handler struct {
	TeamService   service.TeamService
	PlayerService service.PlayerService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R               *gin.Engine
	TeamService     service.TeamService
	PlayerService   service.PlayerService
	TimeoutDuration time.Duration
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	h := &Handler{
		TeamService:   c.TeamService,
		PlayerService: c.PlayerService,
	}

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	c.R.Use(gin.Recovery())
	c.R.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))

	g := c.R.Group("/api/v1")
	{
		g.GET("/teams/:league", h.GetTeams)
		g.GET("/team/:id", h.GetTeamByID)
		g.GET("team/:id/players", h.GetPlayersByTeamID)
	}
}
