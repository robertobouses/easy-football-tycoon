package resume

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) PostPlayerSaleDecision(ctx *gin.Context) {
	var decision struct {
		Accept bool `json:"accept"`
	}

	if err := ctx.BindJSON(&decision); err != nil {
		log.Printf("Error al parsear la decisión de venta: %v", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	player, transferFeeReceived, err := h.app.GetCurrentPlayerSale()
	if err != nil {
		log.Printf("Error al obtener el jugador en venta: %v", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if player == nil {
		log.Printf("No hay jugador disponible para la venta")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No player available"})
		return
	}

	if player.PlayerId == uuid.Nil {
		log.Printf("Error: El PlayerId es vacío o no válido")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El PlayerId es vacío o no válido"})
		return
	}
	log.Printf("Jugador en venta: %+v, Decisión: %v", player, decision.Accept)
	if decision.Accept {
		err := h.app.AcceptPlayerSale(*player)
		if err != nil {
			log.Printf("Error al aceptar la venta del jugador: %v", err)
			ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Sale could not be completed"})
			return
		}
		log.Printf("Venta aceptada, jugador vendido con éxito")
		ctx.JSON(nethttp.StatusOK, gin.H{
			"player":  player.PlayerName,
			"sold by": transferFeeReceived,
			"message": "Player sold successfully"})
	} else {
		h.app.RejectPlayerSale(*player)
		log.Printf("Venta rechazada, jugador no vendido")
		ctx.JSON(nethttp.StatusOK, gin.H{"message": "Player sale rejected"})
	}

	h.app.SetCurrentPlayerSale(nil, nil)
}
