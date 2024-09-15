package strategy

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetStrategy(ctx *gin.Context) {
	player, err := h.app.GetStrategy()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, player)

}
