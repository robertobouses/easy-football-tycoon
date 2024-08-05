package team

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetProspect(ctx *gin.Context) {
	player, err := h.app.GetProspect()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, player)

}
