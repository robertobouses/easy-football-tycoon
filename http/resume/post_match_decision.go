package resume

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/robertobouses/easy-football-tycoon/app"
)

var matchEventStore = make(map[string][]app.EventResult)
var matchScoreStore = make(map[string]app.Match)
var sessionID = "currentMatch"

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

		mu.Lock()
		matchScoreStore[sessionID] = match
		mu.Unlock()

		mu.Lock()
		matchEventStore[sessionID] = events
		mu.Unlock()

		matchIndex = 0

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Iniciando partido contra %v. Haz clic para ver el siguiente evento.", match.RivalName),
		})
		return

	case "simulate":
		match, err = h.app.ProcessMatchSimulation()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al simular el partido"})
			return
		}
		h.sendSimulatedMatchResponse(ctx, match)
		return

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Decisión inválida. Elige 'play' o 'simulate'."})
		return
	}
}
func (h Handler) sendSimulatedMatchResponse(ctx *gin.Context, match app.Match) {
	if match.MatchDayNumber != 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"message":        "Simulación de partido completa",
			"matchDayNumber": match.MatchDayNumber,
			"venue":          match.HomeOrAway,
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
			"result":  match.Result,
			"options": []string{"play", "simulate"},
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Partido simulado correctamente."})
	}
}
