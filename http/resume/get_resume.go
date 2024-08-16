package resume

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetResume(ctx *gin.Context) {
	_, err := h.app.GetResume()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if h.app.GetCurrentSalePlayer() != nil {
		ctx.JSON(nethttp.StatusOK, gin.H{
			"message": "Player on sale",
			"player":  h.app.GetCurrentSalePlayer(),
		})
	} else {
		ctx.JSON(nethttp.StatusOK, gin.H{"message": "Day completed"})
	}
}

func (h Handler) PostSaleDecision(ctx *gin.Context) {
	var decision struct {
		Accept bool `json:"accept"`
	}

	if err := ctx.BindJSON(&decision); err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if decision.Accept {
		err := h.app.AcceptSale(*h.app.GetCurrentSalePlayer())
		if err != nil {
			ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Sale could not be completed"})
			return
		}
		ctx.JSON(nethttp.StatusOK, gin.H{"message": "Player sold successfully"})
	} else {
		h.app.RejectSale(*h.app.GetCurrentSalePlayer())
		ctx.JSON(nethttp.StatusOK, gin.H{"message": "Player sale rejected"})
	}

	h.app.SetCurrentSalePlayer(nil)
}
