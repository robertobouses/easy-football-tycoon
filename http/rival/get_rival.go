package team

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetTeam(ctx *gin.Context) {
	team, err := h.app.GetTeam()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, team)

}
