package app

import (
	"log"
	"math/rand"

	"github.com/google/uuid"
)

func (a *AppService) DistributeAssistsByPossition(goals int) (int, int, int, int) {
	log.Println("goals en DistributeAssistsByPossition", goals)

	totalAssists := rand.Intn(goals + 1)
	log.Println("totalAssists en DistributeAssistsByPossition", totalAssists)

	randomValue := rand.Intn(100)
	log.Println("randomValue en DistributeAssistsByPossition", randomValue)

	var forwardAssists, midfieldAssists, defenderAssists int

	if randomValue < 60 {
		forwardAssists = int(float64(totalAssists) * 0.35)
		midfieldAssists = int(float64(totalAssists) * 0.65)
		defenderAssists = totalAssists - forwardAssists - midfieldAssists
	} else if randomValue < 90 {
		forwardAssists = int(float64(totalAssists) * 0.5)
		midfieldAssists = int(float64(totalAssists) * 0.4)
		defenderAssists = totalAssists - forwardAssists - midfieldAssists
	} else {
		forwardAssists = int(float64(totalAssists) * 0.3)
		midfieldAssists = int(float64(totalAssists) * 0.4)
		defenderAssists = totalAssists - forwardAssists - midfieldAssists
	}
	if forwardAssists+midfieldAssists+defenderAssists > goals {
		forwardAssists = 1
		midfieldAssists = goals - 1
		defenderAssists = 0
	}
	log.Println("forwardAssists, midfieldAssists, defenderAssists, totalAssists en DistributeAssistsByPossition",
		forwardAssists, midfieldAssists, defenderAssists, totalAssists)
	return forwardAssists, midfieldAssists, defenderAssists, totalAssists
}

func (a *AppService) DistributeAssistsByPlayer(lineup []Lineup, forwardAssists, midfieldAssists, defenderAssists, totalAssists int) map[uuid.UUID]int {
	assistsByPlayer := make(map[uuid.UUID]int)

	forwards := filterPlayersByPosition(lineup, "forward")
	midfielders := filterPlayersByPosition(lineup, "midfielder")
	defenders := filterPlayersByPosition(lineup, "defender")

	forwardChancesByPlayer := a.DistributeChances(forwards, forwardAssists)
	for k, v := range forwardChancesByPlayer {
		assistsByPlayer[k] = v
	}

	midfieldChancesByPlayer := a.DistributeChances(midfielders, midfieldAssists)
	for k, v := range midfieldChancesByPlayer {
		assistsByPlayer[k] = v
	}

	defenderChancesByPlayer := a.DistributeChances(defenders, defenderAssists)
	for k, v := range defenderChancesByPlayer {
		assistsByPlayer[k] = v
	}

	log.Println("assistsByPlayer en DistributeAssistsByPlayer", assistsByPlayer)
	return assistsByPlayer
}
