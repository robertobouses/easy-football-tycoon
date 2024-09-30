package analytics

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/easy-football-tycoon/app"
)

type PostAnalyticsRequest struct {
	Training        int `json:"training"`
	Finances        int `json:"finances"`
	Scouting        int `json:"scouting"`
	Physiotherapy   int `json:"physiotherapy"`
	Trust           int `json:"trust"`
	StadiumCapacity int `json:"stadiumcapacity"`
}

func (h Handler) PostAnalytics(c *gin.Context) {
	var req PostAnalyticsRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostAnalytics] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	analytics := app.Analytics{
		Training:        req.Training,
		Finances:        req.Finances,
		Scouting:        req.Scouting,
		Physiotherapy:   req.Physiotherapy,
		Trust:           req.Trust,
		StadiumCapacity: req.StadiumCapacity,
	}

	job := "movimientos iniciales"

	err := h.app.PostAnalytics(analytics, job)
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "analytics insertado correctamente"})
}
