package signings

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

type PostAutoPlayerGeneratorRequest struct {
	NumberOfPlayersToGenerate int `json:"numberofplayertogenerate"`
}

func (h Handler) PostAutoPlayerGenerator(c *gin.Context) {
	var req PostAutoPlayerGeneratorRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostAutoPlayerGeneratorRequest] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.Printf("Número de jugadores a generar: %d", req.NumberOfPlayersToGenerate)

	players, err := h.app.RunAutoPlayerGenerator(req.NumberOfPlayersToGenerate)
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{
		"mensaje":             "Jugadores generados correctamente",
		"jugadores_generados": players,
	})
}
