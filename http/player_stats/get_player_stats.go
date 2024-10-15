package stats

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetPlayerStats(ctx *gin.Context) {
	player, err := h.app.GetPlayerStats()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, player)

}
