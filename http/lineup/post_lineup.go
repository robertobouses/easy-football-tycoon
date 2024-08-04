package lineup

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/easy-football-tycoon/app"
)

type PostLineupRequest struct {
	PlayerID string `json:"playerid"`
}

func (h Handler) PostLineup(c *gin.Context) {
	var req PostLineupRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostLineup] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	lineup := app.Lineup{
		PlayerId: req.PlayerID,
	}
	err := h.app.PostLineup(lineup)

	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "jugador alineado correctamente"})
}
