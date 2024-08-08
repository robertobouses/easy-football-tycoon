package rival

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robertobouses/easy-football-tycoon/app"
)

type PostRivalRequest struct {
	RivalId   uuid.UUID `json:"teamid"`
	RivalName string    `json:"rivalname"`
	Technique int       `json:"technique"`
	Mental    int       `json:"mental"`
	Physique  int       `json:"physique"`
}

func (h Handler) PostRival(c *gin.Context) {
	var req PostRivalRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostRival] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	rival := app.Rival{
		RivalId:   req.RivalId,
		RivalName: req.RivalName,
		Technique: req.Technique,
		Mental:    req.Mental,
		Physique:  req.Physique,
	}
	err := h.app.PostRival(rival)

	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "equipo rival insertado correctamente"})
}
