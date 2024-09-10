package resume

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h Handler) PostStaffSaleDecision(ctx *gin.Context) {
	var decision struct {
		Accept bool `json:"accept"`
	}

	if err := ctx.BindJSON(&decision); err != nil {
		log.Printf("Error al parsear la decisión de venta: %v", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	staff, transferFeeReceived, err := h.app.GetCurrentStaffSale()
	if err != nil {
		log.Printf("Error al obtener el jugador en venta: %v", err)
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if staff == nil {
		log.Printf("No hay jugador disponible para la venta")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No staff available"})
		return
	}

	if staff.StaffId == uuid.Nil {
		log.Printf("Error: El StaffId es vacío o no válido")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El StaffId es vacío o no válido"})
		return
	}
	log.Printf("Jugador en venta: %+v, Decisión: %v", staff, decision.Accept)
	if decision.Accept {
		err := h.app.AcceptStaffSale(*staff)
		if err != nil {
			log.Printf("Error al aceptar la venta del jugador: %v", err)
			ctx.JSON(nethttp.StatusInternalServerError, gin.H{"error": "Sale could not be completed"})
			return
		}
		log.Printf("Venta aceptada, jugador vendido con éxito")
		ctx.JSON(nethttp.StatusOK, gin.H{
			"staff":   staff.StaffName,
			"sold by": transferFeeReceived,
			"message": "Staff sold successfully"})
	} else {
		h.app.RejectStaffSale(*staff)
		log.Printf("Venta rechazada, jugador no vendido")
		ctx.JSON(nethttp.StatusOK, gin.H{"message": "Staff sale rejected"})
	}

	h.app.SetCurrentStaffSale(nil, nil)
}
