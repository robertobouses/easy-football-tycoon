package app

import (
	"log"
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

	strategy, err := a.strategyRepo.GetStrategy()
	if err != nil {
		log.Println("Error obteniendo la estrategia:", err)
		return
	}

	forwardChances, midfieldChances, defenderChances := a.DistributeChancesByStrategy(strategy, teamChances)
	log.Println("-----forwardChances, midfieldChances, defenderChances en CalculatePlayerStats después función DistributeChancesByStrategy", forwardChances, midfieldChances, defenderChances)
	forwardChances, midfieldChances, defenderChances = a.ModifyChancesForBuildUpPlay(strategy, forwardChances, midfieldChances, defenderChances)
	log.Println("-----forwardChances, midfieldChances, defenderChances en CalculatePlayerStats después función ModifyChancesForBuildUpPlay", forwardChances, midfieldChances, defenderChances)
	chancesByPlayer := a.DistributeChancesToPlayers(lineup, forwardChances, midfieldChances, defenderChances, teamChances)
	log.Println("____mapa chancesByPlayer", chancesByPlayer)

	forwardAssists, midfieldAssists, defenderAssists, totalAssists := a.DistributeAssistsByPossition(lastMatch.Team.Goals)
	assistsByPlayer := a.DistributeAssistsByPlayer(lineup, forwardAssists, midfieldAssists, defenderAssists, totalAssists)
	log.Println("-----forwardChances, midfieldChances, defenderChances en CalculatePlayerStats después función DistributeAssistsByPlayer", forwardAssists, midfieldAssists, defenderAssists, totalAssists)
	log.Println("____mapa assistsByPlayer", assistsByPlayer)

	forwardGoals, midfieldGoals, defenderGoals, totalGoals := a.DistributeGoalsByPossition(lastMatch.Team.Goals)
	log.Println("-----forwardChances, midfieldChances, defenderChances en CalculatePlayerStats después función DistributeGoalsByPossition", forwardGoals, midfieldGoals, defenderGoals, totalGoals)
	goalsByPlayer := a.DistributeGoalsByPlayer(lineup, forwardGoals, midfieldGoals, defenderGoals, totalGoals)
	log.Println("____mapa goalsByPlayer", goalsByPlayer)

	for _, player := range lineup {
		playerId := player.PlayerId
		appearances := 1

		chances := chancesByPlayer[playerId]
		assists := assistsByPlayer[playerId]
		goals := goalsByPlayer[playerId]

		log.Println("----playerId en CalculatePlayerStats", playerId)
		log.Println("----chances en CalculatePlayerStats", chances)
		log.Println("----assists en CalculatePlayerStats", assists)
		log.Println("----goals en CalculatePlayerStats", goals)

		err := a.statsRepo.UpdatePlayerStats(playerId, appearances, 0, 0, 0, 0, assists, chances, goals, 0, 0)
		if err != nil {
			log.Println("Error actualizando las estadísticas del jugador:", err)
		}
	}
}
