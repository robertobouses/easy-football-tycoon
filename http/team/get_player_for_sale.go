package team

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetPlayerForSale(ctx *gin.Context) {
	player, err := h.app.Sale()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(nethttp.StatusOK, gin.H{"player": player})
}

func (h Handler) DecideSale(ctx *gin.Context) {
	var request struct {
		Accept bool `json:"accept"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if request.Accept {
		h.app.ProcessSaleDecision(player)
		ctx.JSON(nethttp.StatusOK, gin.H{"message": "Venta aceptada"})
	} else {

		ctx.JSON(nethttp.StatusOK, gin.H{"message": "Venta rechazada"})
	}
}
