package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/new-pop-corn/internal/pkg/model/apperrors"
)

func (h *Handler) GetPlayersByTeamID(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	fmt.Println("id", id)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	teams, err := h.PlayerService.GetPlayersByTeamID(ctx, uint(id))
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	c.JSON(200, teams)
}
