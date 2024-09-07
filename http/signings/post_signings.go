package signings

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robertobouses/easy-football-tycoon/app"
)

type PostSigningsRequest struct {
	SigningsId   uuid.UUID `json:"signingsid"`
	SigningsName string    `json:"signingsname"`
	Position     string    `json:"position"`
	Age          int       `json:"age"`
	Fee          int       `json:"fee"`
	Salary       int       `json:"salary"`
	Technique    int       `json:"technique"`
	Mental       int       `json:"mental"`
	Physique     int       `json:"physique"`
	InjuryDays   int       `json:"injurydays"`
	Rarity       int       `json:"rarity"`
}

func (h Handler) PostSignings(c *gin.Context) {
	var req PostSigningsRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostSignings] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	signings := app.Signings{
		SigningsId:   req.SigningsId,
		SigningsName: req.SigningsName,
		Position:     req.Position,
		Age:          req.Age,
		Fee:          req.Fee,
		Salary:       req.Salary,
		Technique:    req.Technique,
		Mental:       req.Mental,
		Physique:     req.Physique,
		InjuryDays:   req.InjuryDays,
		Rarity:       req.Rarity,
	}

	err := h.app.PostSignings(signings)
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "signings insertado correctamente"})
}
