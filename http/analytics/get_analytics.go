package prospect

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetAnalytics(ctx *gin.Context) {
	analytics, err := h.app.GetAnalytics()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, analytics)

}
