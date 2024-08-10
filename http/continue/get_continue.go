package prospect

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetContinue(ctx *gin.Context) {
	player, err := h.app.GetContinue()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, player)

}
