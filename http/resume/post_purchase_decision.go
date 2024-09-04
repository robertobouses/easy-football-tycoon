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

	prospect, err := h.app.GetCurrentProspect()
	log.Println("prospect en PostPurchaseDecision HTTP", prospect)
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
