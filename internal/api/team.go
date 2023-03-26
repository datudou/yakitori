package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/new-pop-corn/internal/model/apperrors"
)

func (h *Handler) getTeams(c *gin.Context) {
	ctx := c.Request.Context()
	league := c.Param("league")

	teams, err := h.Services.TeamService.GetTeams(ctx, league)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, teams)
}

func (h *Handler) getTeamByID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	teams, err := h.Services.TeamService.GetTeamByID(ctx, uint(id))
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, teams)
}
