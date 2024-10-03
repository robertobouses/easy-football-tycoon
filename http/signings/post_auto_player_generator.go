package signings

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

type PostAutoPlayerGeneratorRequest struct {
	numberOfPlayersToGenerate int `json:"numberofplayertogenerate"`
}

func (h Handler) PostAutoPlayerGenerator(c *gin.Context) {
	var req PostAutoPlayerGeneratorRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostAutoPlayerGeneratorRequest] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	players, err := h.app.RunAutoPlayerGenerator(req.numberOfPlayersToGenerate)
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{
		"mensaje":             "Jugadores generados correctamente",
		"jugadores_generados": players,
	})
}
