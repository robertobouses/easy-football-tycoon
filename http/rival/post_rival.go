package team

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/easy-football-tycoon/app"
)

type PostTeamRequest struct {
	TeamId    string `json:"teamid"`
	TeamName  string `json:"teamname"`
	Technique int    `json:"technique"`
	Mental    int    `json:"mental"`
	Physique  int    `json:"physique"`
}

func (h Handler) PostRival(c *gin.Context) {
	var req PostTeamRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostTeam] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	team := app.Rival{
		TeamId:    req.TeamId,
		TeamName:  req.TeamName,
		Technique: req.Technique,
		Mental:    req.Mental,
		Physique:  req.Physique,
	}
	err := h.app.PostRival(team)

	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "equipo rival insertado correctamente"})
}
