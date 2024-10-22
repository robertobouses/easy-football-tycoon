package resume

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var mu sync.Mutex
var matchIndex int

func (h Handler) PostMatchNextEvent(ctx *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	events, exists := matchEventStore[sessionID]
	if !exists || len(events) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron eventos para este partido."})
		return
	}

	match, exists := matchScoreStore[sessionID]
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se encontró el estado del partido."})
		return
	}

	if matchIndex >= len(events) {
		ctx.JSON(http.StatusOK, gin.H{

			"match": gin.H{
				"your_team": gin.H{
					"ball_possession": match.Team.BallPossession,
					"scoring_chances": match.Team.ScoringChances,
					"goals":           match.Team.Goals,
				},
				"rival": gin.H{
					"rival_name":      match.RivalName,
					"ball_possession": match.Rival.BallPossession,
					"scoring_chances": match.Rival.ScoringChances,
					"goals":           match.Rival.Goals,
				},
			},
		})
		ctx.JSON(http.StatusOK, gin.H{"message": "Todos los eventos han sido mostrados. Fin del partido"})
		return
	}

	event := events[matchIndex]

	ctx.JSON(http.StatusOK, gin.H{
		"event":  event.Event,
		"minute": event.Minute,
	})
	matchIndex++
}
