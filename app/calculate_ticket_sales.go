package app

import "fmt"

//TODO CAPACIDAD ESTADIO INTRODUCIR

func (a *AppService) CalculateTicketSales(stadiumCapacity, TotalLineupQuality, TotalRivalQuality, currentTrust int, previousMatchScoreChances int, previousMatchResult, HomeOrAway string) (int, error) {
	if HomeOrAway == "Away" {
		return 0, nil
	}

	var previousMatchResultFloat float64

	switch previousMatchResult {
	case "LOSE":
		previousMatchResultFloat = 0.7
	case "DRAW":
		previousMatchResultFloat = 0.8
	case "WIN":
		previousMatchResultFloat = 1.3
	default:
		return 0, fmt.Errorf("Resultado anterior no válido: %s", previousMatchResult)
	}

	attendance := float64(TotalLineupQuality) + (4 * float64(TotalRivalQuality)) + float64(currentTrust) + float64(previousMatchScoreChances)*previousMatchResultFloat

	if attendance > float64(stadiumCapacity) {
		attendance = float64(stadiumCapacity)
	}

	ticketSales := int(attendance)

	return ticketSales, nil
}
