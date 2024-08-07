package prospect

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/easy-football-tycoon/app"
)

type PostProspectRequest struct {
	ProspectId   string `json:"prospectid"`
	ProspectName string `json:"prospectname"`
	Position     string `json:"position"`
	Age          int    `json:"age"`
	Fee          int    `json:"fee"`
	Salary       int    `json:"salary"`
	Technique    int    `json:"technique"`
	Mental       int    `json:"mental"`
	Physique     int    `json:"physique"`
	InjuryDays   int    `json:"injurydays"`
	Lined        bool   `json:"lined"`
	Job          string `json:"job"`
	Rarity       int    `json:"rarity"`
}

func (h Handler) PostProspect(c *gin.Context) {
	var req PostProspectRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostProspect] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	prospect := app.Prospect{
		ProspectId:   req.ProspectId,
		ProspectName: req.ProspectName,
		Position:     req.Position,
		Age:          req.Age,
		Fee:          req.Fee,
		Salary:       req.Salary,
		Technique:    req.Technique,
		Mental:       req.Mental,
		Physique:     req.Physique,
		InjuryDays:   req.InjuryDays,
		Lined:        req.Lined,
		Job:          req.Job,
		Rarity:       req.Rarity,
	}

	err := h.app.PostProspect(prospect)
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "prospect insertado correctamente"})
}
