package api

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/new-pop-corn/internal/model/apperrors"
)

func (h *Handler) getGameLogByGameID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	gamelog, err := h.Services.GameService.GetGameLogByGameID(ctx, uint(id))
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, gamelog)
}

func (h *Handler) getGamesByDate(c *gin.Context) {
	ctx := c.Request.Context()
	date := c.Param("date")
	layout := "2006-01-02"
	t, err := time.Parse(layout, date)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	games, err := h.Services.GameService.GetGamesByDate(ctx, t)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, games)
}
