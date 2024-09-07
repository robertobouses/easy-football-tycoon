package signings

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetSignings(ctx *gin.Context) {
	player, err := h.app.GetSignings()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, player)

}
