package lineup

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robertobouses/easy-football-tycoon/app"
)

type PostLineupRequest struct {
	PlayerID uuid.UUID `json:"playerid"`
}

func (h Handler) PostLineup(c *gin.Context) {
	var req PostLineupRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostLineup] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err := h.app.PostLineup(req.PlayerID)

	if err != nil {
		if err == app.ErrPlayerNotFound {
			c.JSON(nethttp.StatusNotFound, gin.H{"error": "El jugador no existe en la base de datos"})
		} else {
			c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "jugador alineado correctamente"})
}
