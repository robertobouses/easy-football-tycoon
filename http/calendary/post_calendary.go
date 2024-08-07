package calendary

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) PostCalendary(c *gin.Context) {

	err := h.app.PostCalendary()
	if err != nil {
		log.Printf("[PostCalendary] error inserting calendary: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "calendary insertado correctamente"})
}
