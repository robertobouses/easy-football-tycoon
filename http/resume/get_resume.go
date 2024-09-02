package resume

import (
	"log"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) GetResume(ctx *gin.Context) {
	if h.app == nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "App service is not initialized"})
		return
	}

	calendary, err := h.app.GetResume()
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
			"message":            "Player on sale",
			"player":             player,
			"type calendary day": calendary,
		})
		return
	}

	prospect, err := h.app.GetCurrentProspect()
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Error fetching current prospect"})
		return
	}
	log.Println("el prospect en GetResume HHTP, es", prospect)

	if prospect != nil && prospect.ProspectId != uuid.Nil {
		ctx.JSON(nethttp.StatusOK, gin.H{
			"message":            "Prospect on purchase",
			"prospect":           prospect,
			"type calendary day": calendary,
		})
		return
	}

	ctx.JSON(nethttp.StatusOK, gin.H{
		"message":            "Day completed",
		"type calendary day": calendary,
	})
}
