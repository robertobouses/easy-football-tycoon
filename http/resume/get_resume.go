package resume

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetResume(ctx *gin.Context) {
	player, err := h.app.GetResume()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, player)

}
