package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/new-pop-corn/internal/middleware"
	"github.com/new-pop-corn/internal/model/apperrors"
	"github.com/new-pop-corn/internal/service"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
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
	store := persistence.NewInMemoryStore(24 * time.Hour)

	c.R.Use(gin.Recovery())
	c.R.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))

	c.R.GET("/health", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
	})

	g := c.R.Group("/api/v1")
	{
		g.GET("/teams/:league", h.getTeams)
		g.GET("/team/:id", h.getTeamByID)
		g.GET("/team/:id/players", h.getPlayersByTeamID)
		g.GET("/games/:date", h.getGamesByDate)
		g.GET("/game/:id/gamelog", cache.CachePage(store, 24*time.Hour, h.getGameLogByGameID))
	}
}
