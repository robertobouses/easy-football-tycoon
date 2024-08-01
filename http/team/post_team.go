package team

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/easy-football-tycoon/app"
)

type PostTeamRequest struct {
	PlayerName string `json:"playername"`
	Position   string `json:"position"`
	Age        int    `json:"age"`
	Fee        int    `json:"fee"`
	Salary     int    `json:"salary"`
	Technique  int    `json:"technique"`
	Mental     int    `json:"mental"`
	Physique   int    `json:"physique"`
}

func (h Handler) PostTeam(c *gin.Context) {
	var req PostTeamRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostTeam] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	team := app.Team{
		PlayerName: req.PlayerName,
		Position:   req.Position,
		Age:        req.Age,
		Fee:        req.Fee,
		Salary:     req.Salary,
		Technique:  req.Technique,
		Mental:     req.Mental,
		Physique:   req.Physique,
		InjuryDays: 0,
		Lined:      false,
	}
	err := h.app.PostTeam(team)

	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, req)
	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "jugador insertado correctamente"})
}
