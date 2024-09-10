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

	playerOnSale, transferFeeReceived, err := h.app.GetCurrentPlayerSale()
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Error fetching current sale player"})
		return
	}
	if playerOnSale != nil {
		ctx.JSON(nethttp.StatusOK, gin.H{
			"message":               "Player on sale",
			"player":                playerOnSale,
			"type calendary day":    calendary,
			"Transfer fee received": transferFeeReceived,
		})
		return
	}

	playerSigning, err := h.app.GetCurrentPlayerSigning()
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Error fetching current signings"})
		return
	}
	log.Println("el signings en GetResume HHTP, es", playerSigning)

	if playerSigning != nil && playerSigning.SigningsId != uuid.Nil {
		ctx.JSON(nethttp.StatusOK, gin.H{
			"message":            "Player on Signing",
			"signings":           playerSigning,
			"type calendary day": calendary,
		})
		return
	}

	staffSigning, err := h.app.GetCurrentStaffSigning()
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Error fetching current signings"})
		return
	}
	log.Println("el signings en GetResume HHTP, es", staffSigning)

	if staffSigning != nil && staffSigning.StaffId != uuid.Nil {
		ctx.JSON(nethttp.StatusOK, gin.H{
			"message":            "Staff on Signing",
			"signings":           staffSigning,
			"type calendary day": calendary,
		})
		return
	}

	staffOnSale, transferFeeReceived, err := h.app.GetCurrentStaffSale()
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Error fetching current sale player"})
		return
	}
	if staffOnSale != nil {
		ctx.JSON(nethttp.StatusOK, gin.H{
			"message":               "Staff on sale",
			"player":                staffOnSale,
			"type calendary day":    calendary,
			"Transfer fee received": transferFeeReceived,
		})
		return
	}

	injuredPlayer, injuryDays, err := h.app.GetCurrentInjuredPlayer()
	if err != nil {
		ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Error fetching current signings"})
		return
	}
	log.Println("The injured player in GetCurrentInjuredPlayer HTTP is", injuredPlayer)

	if injuredPlayer != nil && injuredPlayer.PlayerId != uuid.Nil {
		ctx.JSON(nethttp.StatusOK, gin.H{
			"message":            "Player injured",
			"player":             injuredPlayer,
			"type calendary day": calendary,
			"injury days":        injuryDays,
		})
		return
	}

	ctx.JSON(nethttp.StatusOK, gin.H{
		"message":            "Day completed",
		"type calendary day": calendary,
	})
}
