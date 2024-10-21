package resume

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/easy-football-tycoon/app"
)

func (h Handler) PostMatchDecision(ctx *gin.Context) {
	var request struct {
		Decision string `json:"decision"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error en la solicitud. Asegúrate de enviar un JSON válido."})
		return
	}

	decision := request.Decision
	fmt.Println("Decision received:", decision)

	var match app.Match
	var events []app.EventResult
	var err error

	switch decision {
	case "play":
		match, events, err = h.app.ProcessMatchPlay()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al jugar el partido"})
			return
		}

	case "simulate":
		match, err = h.app.ProcessMatchSimulation()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al simular el partido"})
			return
		}

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Decisión inválida. Elige 'play' o 'simulate'."})
		return
	}

	h.sendMatchResponse(ctx, match)

	for _, event := range events {
		if event.Event == "" {
			continue
		}

		time.Sleep(1 * time.Second)

		ctx.SSEvent("match_event", gin.H{
			"minute": event.Minute,
			"event":  event.Event,
		})
	}
}
func (h Handler) sendMatchResponse(ctx *gin.Context, match app.Match) {
	if match.MatchDayNumber != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message":        "Match Playing",
			"matchDayNumber": match.MatchDayNumber,
			"venue":          match.HomeOrAway,
			"match": gin.H{
				"your_team": gin.H{
					"ball_possession": match.Team.BallPossession,
					"scoring_chances": match.Team.ScoringChances,
					"goals":           match.Team.Goals,
				},
				"rival": gin.H{
					"rival name":      match.RivalName,
					"ball_possession": match.Rival.BallPossession,
					"scoring_chances": match.Rival.ScoringChances,
					"goals":           match.Rival.Goals,
				},
			},
			"result":  match.Result,
			"options": []string{"play", "simulate"},
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Partido simulado correctamente."})
	}
}
