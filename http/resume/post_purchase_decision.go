package resume

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) PostPurchaseDecision(ctx *gin.Context) {
	var decision struct {
		Accept bool `json:"accept"`
	}

	if err := ctx.BindJSON(&decision); err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	signings, err := h.app.GetCurrentPlayerSigning()
	log.Println("signings en PostPurchaseDecision HTTP", signings)
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if signings == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No signings available"})
		return
	}

	if decision.Accept {
		err := h.app.AcceptPlayerSigning(signings)
		if err != nil {
			ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Could not complete purchase"})
			return
		}
		ctx.JSON(nethttp.StatusOK, gin.H{"message": "Signings purchased successfully"})
	} else {
		h.app.RejectPlayerSigning(signings)
		ctx.JSON(nethttp.StatusOK, gin.H{"message": "Signings purchase rejected"})
	}

	h.app.SetCurrentSigningPlayer(nil)
}
