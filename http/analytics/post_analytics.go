package analytics

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/easy-football-tycoon/app"
)

type PostAnalyticsRequest struct {
	// AnalyticsId   uuid.UUID `json:"analyticsid"`
	Finances      int `json:"finances"`
	Scouting      int `json:"scouting"`
	Physiotherapy int `json:"physiotherapy"`
}

func (h Handler) PostAnalytics(c *gin.Context) {
	var req PostAnalyticsRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostAnalytics] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	analytics := app.Analytics{
		Finances:      req.Finances,
		Scouting:      req.Scouting,
		Physiotherapy: req.Physiotherapy,
	}

	err := h.app.PostAnalytics(analytics)
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "analytics insertado correctamente"})
}
