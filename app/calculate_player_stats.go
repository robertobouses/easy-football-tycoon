package app

import (
	"log"
)

const (
	possessionWeight  = 0.4
	happinessWeight   = 0.2
	familiarityWeight = 0.2
	fitnessWeight     = 0.2
)

func (a *AppService) CalculatePlayerStats() {

	matches, err := a.matchRepo.GetMatches()
	if err != nil {
		log.Println("Error obteniendo los partidos:", err)
		return
	}

	if len(matches) == 0 {
		log.Println("No hay partidos disponibles.")
		return
	}
	lastMatch := matches[len(matches)-1]

	teamChances := lastMatch.Team.ScoringChances

	lineup, err := a.lineupRepo.GetLineup()
	if err != nil {
		log.Println("Error obteniendo la alineación:", err)
		return
	}

	strategy, err := a.GetStrategy()
	if err != nil {
		log.Println("Error obteniendo la estrategia:", err)
		return
	}

	forwardChances, midfieldChances, defenderChances := a.DistributeChancesByStrategy(strategy, teamChances)
	forwardChances, midfieldChances, defenderChances = a.ModifyChancesForBuildUpPlay(strategy, forwardChances, midfieldChances, defenderChances)
	chancesByPlayer := a.DistributeChancesToPlayers(lineup, forwardChances, midfieldChances, defenderChances, teamChances)
	log.Println("____mapa chancesByPlayer", chancesByPlayer)

	forwardAssists, midfieldAssists, defenderAssists, totalAssists := a.DistributeAssistsByPossition(lastMatch.Team.Goals)
	assistsByPlayer := a.DistributeAssistsByPlayer(lineup, forwardAssists, midfieldAssists, defenderAssists, totalAssists)
	log.Println("____mapa assistsByPlayer", assistsByPlayer)

	forwardGoals, midfieldGoals, defenderGoals, totalGoals := a.DistributeGoalsByPossition(lastMatch.Team.Goals)
	goalsByPlayer := a.DistributeGoalsByPlayer(lineup, forwardGoals, midfieldGoals, defenderGoals, totalGoals)
	log.Println("____mapa goalsByPlayer", goalsByPlayer)

	for _, player := range lineup {
		playerId := player.PlayerId
		appearances := 1

		chances := chancesByPlayer[playerId]
		assists := assistsByPlayer[playerId]
		goals := goalsByPlayer[playerId]

		log.Println("playerId en CalculatePlayerStats", playerId)
		log.Println("chances en CalculatePlayerStats", chances)
		log.Println("assists en CalculatePlayerStats", assists)
		log.Println("goals en CalculatePlayerStats", goals)

		err := a.statsRepo.UpdatePlayerStats(playerId, appearances, chances, assists, goals, 0, 0)
		if err != nil {
			log.Println("Error actualizando las estadísticas del jugador:", err)
		}
	}
}

// // Repartir las asistencias basadas en la posición
//

// // Repartir los goles sin asistencia
// func distributeGoalsWithoutAssists(strategy Strategy, remainingGoals int, lineup []Lineup) {
// 	for _, player := range lineup {
// 		// Distribuir aleatoriamente con una probabilidad de 25% de que cualquier jugador anote un gol
// 		if rand.Intn(100) < 25 {
// 			player.Goals += 1
// 			remainingGoals -= 1
// 			if remainingGoals <= 0 {
// 				break
// 			}
// 		}
// 	}

// 	// Distribuir el resto de los goles según la calidad del jugador y posición
// 	if remainingGoals > 0 {
// 		// Asignar goles a delanteros y mediocampistas basados en calidad
// 		for _, player := range lineup {
// 			if remainingGoals <= 0 {
// 				break
// 			}

// 			// Repartir 75% de los goles entre delanteros y mediocampistas
// 			if player.Position == "forward" && rand.Intn(100) < 60 {
// 				player.Goals += 1
// 				remainingGoals -= 1
// 			} else if player.Position == "midfield" && rand.Intn(100) < 40 {
// 				player.Goals += 1
// 				remainingGoals -= 1
// 			}
// 		}
// 	}
// }

// func calculatePlayerRating(player Player, possession float64, goals, assists, rivalChances, teamChances, rivalGoals int, result string) float64 {
// 	playerQuality := player.Technique + player.Mental + player.Physique
// 	rating := float64(playerQuality) * possessionWeight
// 	rating += float64(player.Happiness) * happinessWeight
// 	rating += float64(player.Familiarity) * familiarityWeight
// 	rating += float64(player.Fitness) * fitnessWeight

// 	switch {
// 	case possession > 70:
// 		rating += 0.9
// 	case possession > 60:
// 		rating += 0.6
// 	case possession > 50:
// 		rating += 0.4
// 	}

// 	switch {
// 	case 2*rivalChances < teamChances:
// 		rating += 0.7
// 	case int(1.5*float64(rivalChances)) < teamChances:
// 		rating += 0.5
// 	case rivalChances < teamChances:
// 		rating += 0.2
// 	case rivalChances > teamChances:
// 		rating -= 0.2
// 	case rivalChances > int(1.5*float64(teamChances)):
// 		rating -= 0.5
// 	case rivalChances > 2*teamChances:
// 		rating -= 0.7
// 	default:
// 		rating += 0
// 	}

// 	rating += float64(goals)
// 	rating += float64(assists)

// 	rating = math.Max(0, math.Min(10, rating))

// 	return rating
// }

// // Repartir asistencias y goles
// assists := distributeAssists(teamGoals)
// remainingGoals := teamGoals - assists
// distributeGoalsWithoutAssists(strategy, remainingGoals, lineup)

// player.Rating = calculatePlayerRating(player, teamBallPossession, playerLineup.Goals, playerLineup.Assists, rivalChances, rivalGoals, result)

// // Verificar si es el mejor jugador (MVP)
// if player.Rating > highestRating {
// 	highestRating = player.Rating
// 	mvpPlayer = player
// }

// }

// var highestRating float64 = 0
// 	var mvpPlayer *Player

// 	for _, playerLineup := range lineup {
// 		player, err := a.teamRepo.GetPlayerByPlayerId(playerLineup.PlayerId)
// 		if err != nil {
// 			log.Println("Error obteniendo al jugador:", err)
// 			continue
// 		}

// 		if mvpPlayer != nil {
// 			mvpPlayer.MVPs += 1
// 			log.Printf("El MVP del partido es: %s con una nota de %.2f", mvpPlayer.Name, mvpPlayer.Rating)
// 		}
//
// teamBallPossession := lastMatch.Team.BallPossession
// 	teamGoals := lastMatch.Team.Goals
// 	rivalChances := lastMatch.Rival.ScoringChances
// 	rivalGoals := lastMatch.Rival.Goals
// 	result := lastMatch.Result
