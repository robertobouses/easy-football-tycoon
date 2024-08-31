package resume

import (
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetResume(ctx *gin.Context) {
	if h.app == nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "App service is not initialized"})
		return
	}

	_, err := h.app.GetResume()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	player, err := h.app.GetCurrentSalePlayer()
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Error fetching current sale player"})
		return
	}
	if player != nil {
		ctx.JSON(nethttp.StatusOK, gin.H{
			"message": "Player on sale",
			"player":  player,
		})
		return
	}

	prospect, err := h.app.GetCurrentProspect()
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Error fetching current prospect"})
		return
	}
	if prospect != nil {
		ctx.JSON(nethttp.StatusOK, gin.H{
			"message": "Prospect on purchase",
			"player":  prospect,
		})
		return
	}

	ctx.JSON(nethttp.StatusOK, gin.H{"message": "Day completed"})
}

func (h Handler) PostPurchaseDecision(ctx *gin.Context) {
	var decision struct {
		Accept bool `json:"accept"`
	}

	if err := ctx.BindJSON(&decision); err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	prospect, err := h.app.GetCurrentProspect()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if prospect == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No prospect available"})
		return
	}

	if decision.Accept {
		err := h.app.AcceptPurchase(prospect)
		if err != nil {
			ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Could not complete purchase"})
			return
		}
		ctx.JSON(nethttp.StatusOK, gin.H{"message": "Prospect purchased successfully"})
	} else {
		h.app.RejectPurchase(prospect)
		ctx.JSON(nethttp.StatusOK, gin.H{"message": "Prospect purchase rejected"})
	}

	h.app.SetCurrentProspect(nil)
}

func (h Handler) PostSaleDecision(ctx *gin.Context) {
	var decision struct {
		Accept bool `json:"accept"`
	}

	if err := ctx.BindJSON(&decision); err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	player, err := h.app.GetCurrentSalePlayer()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if player == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No player available"})
		return
	}
	if decision.Accept {
		err := h.app.AcceptSale(*player)
		if err != nil {
			ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Sale could not be completed"})
			return
		}
		ctx.JSON(nethttp.StatusOK, gin.H{"message": "Player sold successfully"})
	} else {
		h.app.RejectSale(*player)
		ctx.JSON(nethttp.StatusOK, gin.H{"message": "Player sale rejected"})
	}

	h.app.SetCurrentSalePlayer(nil)
}
