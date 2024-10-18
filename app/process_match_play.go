package app

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (a *AppService) ProcessMatchPlay() (Match, []EventResult, error) {

	lineup, err := a.lineupRepo.GetLineup()
	if err != nil {
		log.Println("error al obtener lineup", err)
		return Match{}, []EventResult{}, err
	}

	rival, homeOrAwayString, matchDayNumber, err := a.GetCurrentRival()
	if err != nil {
		log.Println("Error al obtener el rival:", err)
		return Match{}, []EventResult{}, err
	}

	var homeOrAway HomeOrAway
	switch homeOrAwayString {
	case "home":
		homeOrAway = Home
	case "away":
		homeOrAway = Away
	default:
		return Match{}, []EventResult{}, errors.New("valor inválido para homeOrAway")
	}
	log.Printf("Rival seleccionado: %s\n", rival.RivalName)
	log.Printf("Día del partido: %d\n", matchDayNumber)

	rivalLineup := a.SimulateRivalLineup(rival)

	numberOfMatchEvents, err := a.CalculateNumberOfMatchEvents()
	if err != nil {
		log.Println("error al obtener numberOfMatchEvents", err)
		return Match{}, []EventResult{}, err
	}

	numberOfLineupEvents, numberOfRivalEvents, err := DistributeMatchEvents(lineup, rivalLineup, numberOfMatchEvents, homeOrAwayString)
	if err != nil {
		log.Println("error al distribuir numberOfMatchEvents", err)
		return Match{}, []EventResult{}, err
	}

	lineupResults, rivalResults, lineupScoreChances, rivalScoreChances, lineupGoals, rivalGoals := a.GenerateEvents(lineup, rivalLineup, numberOfLineupEvents, numberOfRivalEvents)
	allEvents := append(lineupResults, rivalResults...)
	totalTechnique := lineup[0].TotalTechnique
	totalMental := lineup[0].TotalMental
	totalPhysique := lineup[0].TotalPhysique

	strategy, err := a.GetStrategy()
	if err != nil {
		log.Println("Error al obtener la estrategia:", err)
		return Match{}, []EventResult{}, fmt.Errorf("error al obtener la estrategia: %w", err)
	}
	log.Printf("Estrategia recuperada: %+v\n", strategy)

	resultOfStrategy, err := a.CalculateResultOfStrategy(lineup, strategy.Formation, strategy.PlayingStyle, strategy.GameTempo, strategy.PassingStyle, strategy.DefensivePositioning, strategy.BuildUpPlay, strategy.AttackFocus, strategy.KeyPlayerUsage)
	if err != nil {
		log.Println("Error al calcular el resultado de la estrategia:", err)
		return Match{}, []EventResult{}, fmt.Errorf("error al calcular el resultado de la estrategia: %w", err)
	}

	totalPhysique = totalPhysique + int(resultOfStrategy["teamPhysique"])

	lineupTotalQuality, rivalTotalQuality, allQuality, err := CalculateTotalQuality(totalTechnique, totalMental, totalPhysique, rival.Technique, rival.Mental, rival.Physique)
	if err != nil {
		log.Println("Error al calcular la calidad total:", err)
		return Match{}, []EventResult{}, err
	}
	log.Printf("Calidad total: jugador %d, rival %d, calidad total %d\n", lineupTotalQuality, rivalTotalQuality, allQuality)
	lineupPercentagePossession, rivalPercentagePossession, err := a.CalculateBallPossession(totalTechnique, totalMental, lineupTotalQuality, rivalTotalQuality, allQuality, homeOrAwayString, resultOfStrategy["teamPossession"])

	var result string
	if lineupGoals > rivalGoals {
		result = "WIN"
	} else if lineupGoals < rivalGoals {
		result = "LOSE"
	} else {
		result = "DRAW"
	}
	match := Match{
		MatchId:        uuid.New(),
		MatchDayNumber: matchDayNumber,
		HomeOrAway:     homeOrAway,
		RivalName:      rival.RivalName,
		Team: TeamStats{
			BallPossession: lineupPercentagePossession,
			ScoringChances: lineupScoreChances,
			Goals:          lineupGoals,
		},
		Rival: TeamStats{
			BallPossession: rivalPercentagePossession,
			ScoringChances: rivalScoreChances,
			Goals:          rivalGoals,
		},
		Result: result,
	}

	return match, allEvents, nil
}
