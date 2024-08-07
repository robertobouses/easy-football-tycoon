package calendary

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetCalendary(ctx *gin.Context) {
	calendary, err := h.app.GetCalendary()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, calendary)

}
