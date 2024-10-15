package app

import (
	"log"
	"math/rand"

	"github.com/google/uuid"
)

func (a *AppService) DistributeGoalsByPossition(goals int) (int, int, int, int) {
	totalGoals := rand.Intn(goals + 1)
	if totalGoals >= goals {
		totalGoals = goals
	}
	var forwardGoals, midfieldGoals, defenderGoals int
	randomValue := rand.Intn(100)
	if randomValue < 60 {
		forwardGoals = int(float64(totalGoals) * 0.35)
		midfieldGoals = int(float64(totalGoals) * 0.65)
		defenderGoals = 0
	} else if randomValue < 90 {
		forwardGoals = int(float64(totalGoals) * 0.5)
		midfieldGoals = int(float64(totalGoals) * 0.4)
		defenderGoals = int(float64(totalGoals) * 0.1)
	} else {
		forwardGoals = int(float64(totalGoals) * 0.3)
		midfieldGoals = int(float64(totalGoals) * 0.4)
		defenderGoals = int(float64(totalGoals) * 0.3)
	}

	totalDistribuidos := forwardGoals + midfieldGoals + defenderGoals
	if totalDistribuidos > totalGoals {
		forwardGoals = goals
		midfieldGoals = 0
		defenderGoals = 0
	}

	log.Println("forwardGoals, midfieldGoals, defenderGoals, totalGoals en DistributeGoalsByPossition",
		forwardGoals, midfieldGoals, defenderGoals, totalGoals)

	return forwardGoals, midfieldGoals, defenderGoals, totalGoals
}

func (a *AppService) DistributeGoalsByPlayer(lineup []Lineup, forwardGoals, midfieldGoals, defenderGoals, totalGoals int) map[uuid.UUID]int {
	GoalsByPlayer := make(map[uuid.UUID]int)

	forwards := filterPlayersByPosition(lineup, "forward")
	midfielders := filterPlayersByPosition(lineup, "midfielder")
	defenders := filterPlayersByPosition(lineup, "defender")
	log.Println("defenders on DistributeGoalsByPlayer", defenders)

	forwardChancesByPlayer := a.DistributeChances(forwards, forwardGoals)
	for k, v := range forwardChancesByPlayer {
		GoalsByPlayer[k] = v
	}

	midfieldChancesByPlayer := a.DistributeChances(midfielders, midfieldGoals)
	for k, v := range midfieldChancesByPlayer {
		GoalsByPlayer[k] = v
	}

	defenderChancesByPlayer := a.DistributeChances(defenders, defenderGoals)
	for k, v := range defenderChancesByPlayer {
		GoalsByPlayer[k] = v
	}

	log.Println("GoalsByPlayer en DistributeGoalsByPlayer", GoalsByPlayer)
	return GoalsByPlayer
}
