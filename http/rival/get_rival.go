package rival

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetRival(ctx *gin.Context) {
	team, err := h.app.GetRival()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, team)

}
