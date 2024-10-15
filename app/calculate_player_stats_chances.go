package app

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const (
	possessionWeight  = 0.4
	happinessWeight   = 0.2
	familiarityWeight = 0.2
	fitnessWeight     = 0.2
)

func (a *AppService) DistributeChancesByStrategy(strategy Strategy, teamChances int) (int, int, int) {
	var forwardChances, midfieldChances, defenderChances int
	switch strategy.PassingStyle {
	case "possession":
		forwardChances = int(0.4 * float64(teamChances))
		midfieldChances = int(0.4 * float64(teamChances))
		defenderChances = int(0.2 * float64(teamChances))
	case "directplay":
		forwardChances = int(0.7 * float64(teamChances))
		midfieldChances = int(0.2 * float64(teamChances))
		defenderChances = int(0.1 * float64(teamChances))
	case "lowblock":
		forwardChances = int(0.8 * float64(teamChances))
		midfieldChances = int(0.15 * float64(teamChances))
		defenderChances = int(0.05 * float64(teamChances))
	default:
		forwardChances = int(0.6 * float64(teamChances))
		midfieldChances = int(0.3 * float64(teamChances))
		defenderChances = int(0.1 * float64(teamChances))
	}
	return forwardChances, midfieldChances, defenderChances
	return 10, 6, 2
}

func (a *AppService) ModifyChancesForBuildUpPlay(strategy Strategy, forwardChances, midfieldChances, defenderChances int) (int, int, int) {
	switch strategy.BuildUpPlay {
	case "play_from_back":
		midfieldChances += int(0.1 * float64(midfieldChances))
		defenderChances += int(0.05 * float64(defenderChances))
		forwardChances -= int(0.15 * float64(forwardChances))
	case "long_clearance":
		forwardChances += int(0.07 * float64(forwardChances))
		defenderChances -= int(0.07 * float64(defenderChances))
	}
	return forwardChances, midfieldChances, defenderChances
}

func (a *AppService) DistributeChancesToPlayers(lineup []Lineup, forwardChances, midfieldChances, defenderChances, totalChances int) map[uuid.UUID]int {
	chancesByPlayer := make(map[uuid.UUID]int)

	forwards := filterPlayersByPosition(lineup, "forward")
	midfielders := filterPlayersByPosition(lineup, "midfielder")
	defenders := filterPlayersByPosition(lineup, "defender")

	forwardChancesByPlayer := a.DistributeChances(forwards, forwardChances)
	for k, v := range forwardChancesByPlayer {
		chancesByPlayer[k] = v
	}

	midfieldChancesByPlayer := a.DistributeChances(midfielders, midfieldChances)
	for k, v := range midfieldChancesByPlayer {
		chancesByPlayer[k] = v
	}

	defenderChancesByPlayer := a.DistributeChances(defenders, defenderChances)
	for k, v := range defenderChancesByPlayer {
		chancesByPlayer[k] = v
	}

	return chancesByPlayer
}

func (a *AppService) DistributeChances(players []Lineup, totalChances int) map[uuid.UUID]int {
	chancesByPlayer := make(map[uuid.UUID]int)
	if len(players) == 0 {
		return chancesByPlayer
	}

	rand.Seed(time.Now().UnixNano())

	totalWeight := 0.0
	for _, player := range players {
		playerWeight := a.calculatePlayerWeight(player)
		totalWeight += playerWeight
	}

	for _, player := range players {
		playerWeight := a.calculatePlayerWeight(player)

		playerChances := int((playerWeight / totalWeight) * float64(totalChances))

		randomFactor := rand.Float64() * 0.45
		playerChances = int(float64(playerChances) * (1 + randomFactor))

		chancesByPlayer[player.PlayerId] = playerChances
	}

	log.Println("chancesByPlayer en DistributeChances", chancesByPlayer)
	return chancesByPlayer
}

func filterPlayersByPosition(lineup []Lineup, position string) []Lineup {
	var playersInPosition []Lineup
	for _, player := range lineup {
		if player.Position == position {
			playersInPosition = append(playersInPosition, player)
		}
	}
	return playersInPosition
}
func (a *AppService) calculatePlayerWeight(lineupPlayer Lineup) float64 {
	player, err := a.teamRepo.GetPlayerByPlayerId(lineupPlayer.PlayerId)
	if err != nil {
		log.Println("Error al obtener jugador por ID:", err)
		return 0
	}

	technique := float64(player.Technique)
	happiness := float64(player.Happiness)
	familiarity := float64(player.Familiarity)
	fitness := float64(player.Fitness)

	playerWeight := (technique * possessionWeight) +
		(happiness * happinessWeight) +
		(familiarity * familiarityWeight) +
		(fitness * fitnessWeight)

	return playerWeight
}
