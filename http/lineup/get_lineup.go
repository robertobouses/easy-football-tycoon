package lineup

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetLineup(ctx *gin.Context) {
	lineup, err := h.App.GetLineup()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(nethttp.StatusOK, lineup)

}
