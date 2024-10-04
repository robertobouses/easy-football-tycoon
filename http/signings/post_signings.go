package signings

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robertobouses/easy-football-tycoon/app/signings"
)

type PostSigningsRequest struct {
	SigningsId  uuid.UUID `json:"signingsid"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Nationality string    `json:"nationality"`
	Position    string    `json:"position"`
	Age         int       `json:"age"`
	Fee         int       `json:"fee"`
	Salary      int       `json:"salary"`
	Technique   int       `json:"technique"`
	Mental      int       `json:"mental"`
	Physique    int       `json:"physique"`
	InjuryDays  int       `json:"injurydays"`
	Rarity      int       `json:"rarity"`
}

func (h Handler) PostSignings(c *gin.Context) {
	var req PostSigningsRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostSignings] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	signing := signings.Signings{
		SigningsId:  req.SigningsId,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Nationality: req.Nationality,
		Position:    req.Position,
		Age:         req.Age,
		Fee:         req.Fee,
		Salary:      req.Salary,
		Technique:   req.Technique,
		Mental:      req.Mental,
		Physique:    req.Physique,
		InjuryDays:  req.InjuryDays,
		Rarity:      req.Rarity,
	}

	err := h.app.PostSignings(signing)
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "signings insertado correctamente"})
}
