package app

import (
	"errors"
	"fmt"
)

func (a *AppService) ResultOfStrategy(lineup []Lineup, formation, playingStyle, gameTempo, passingStyle, defensivePositioning, buildUpPlay, attackFocus, keyPlayerUsage string) (result map[string]float64, err error) {

	result = make(map[string]float64)

	teamPossessionFormation, teamChancesFormation, rivalChancesFormation, err := a.CalculatePossessionChancesByFormation(lineup, formation)
	if err != nil {
		return nil, fmt.Errorf("Error en la formación: %v", err)
	}

	teamPossessionStyle, teamChancesStyle, rivalChancesStyle, physiqueStyle, err := a.CalculatePossessionChancesByPlayingStyle(lineup, playingStyle)
	if err != nil {
		return nil, fmt.Errorf("Error en el estilo de juego: %v", err)
	}

	teamPossessionTempo, teamChancesTempo, rivalChancesTempo, physiqueTempo, err := a.CalculatePossessionChancesByGameTempo(gameTempo)
	if err != nil {
		return nil, fmt.Errorf("Error en el tempo del juego: %v", err)
	}

	teamPossessionPassing, rivalChancesPassing, err := a.CalculatePossessionChancesByPassingStyle(passingStyle)
	if err != nil {
		return nil, fmt.Errorf("Error en el estilo de pase: %v", err)
	}

	rivalChancesDefensivePositioning, physiqueDefensive, err := a.CalculateRivalChancesByDefensivePositioning(lineup, defensivePositioning)
	if err != nil {
		return nil, fmt.Errorf("Error en el posicionamiento defensivo: %v", err)
	}

	teamPossessionBuildUpPlay, err := a.CalculatePossessionByBuildUpPlay(buildUpPlay)
	if err != nil {
		return nil, fmt.Errorf("Error en el estilo de build-up play: %v", err)
	}

	rivalChancesAttackFocus, err := a.CalculateRivalChancesByAttackFocus(lineup, attackFocus)
	if err != nil {
		return nil, fmt.Errorf("Error en el enfoque de ataque: %v", err)
	}

	teamPossessionKeyPlayer, teamChancesKeyPlayer, err := a.CalculateRivalChancesByKeyPlayerUsage(lineup, keyPlayerUsage)
	if err != nil {
		return nil, fmt.Errorf("Error en el uso del jugador clave: %v", err)
	}

	totalTeamPossession := (teamPossessionFormation + teamPossessionStyle + teamPossessionTempo + teamPossessionPassing + teamPossessionBuildUpPlay + teamPossessionKeyPlayer) / 6
	totalTeamChances := (teamChancesFormation + teamChancesStyle + teamChancesTempo + teamChancesKeyPlayer) / 4
	totalRivalChances := (rivalChancesFormation + rivalChancesStyle + rivalChancesTempo + rivalChancesPassing + rivalChancesDefensivePositioning + rivalChancesAttackFocus) / 6

	totalPhysique := physiqueStyle + physiqueTempo + physiqueDefensive

	result["teamPossession"] = totalTeamPossession
	result["teamChances"] = totalTeamChances
	result["rivalChances"] = totalRivalChances
	result["teamPhysique"] = float64(totalPhysique)

	return result, nil
}

func (a *AppService) CalculatePossessionChancesByFormation(lineup []Lineup, formation string) (teamPossession, teamChances, rivalChances float64, err error) {

	totalDefendersQuality, err := getTwoBestPlayers(lineup, "defender")
	totalMidfieldersQuality, err := getTwoBestPlayers(lineup, "midfielder")
	totalForwardersQuality, err := getTwoBestPlayers(lineup, "forwarder")

	switch formation {
	case "4-4-2":
		if totalForwardersQuality >= 550 {
			return 0.9, 1.2, 1, nil
		} else if totalForwardersQuality >= 360 {
			return 0.9, 1, 1, nil
		} else {
			return 0.8, 1, 1, nil
		}

	case "4-3-3":
		if totalMidfieldersQuality >= 510 {
			return 0.9, 1.2, 1.1, nil
		} else {
			return 0.8, 1.2, 1.1, nil
		}

	case "4-5-1":
		if totalMidfieldersQuality >= 540 {
			return 1.4, 0.7, 0.7, nil
		} else if totalMidfieldersQuality >= 380 {
			return 1.3, 0.7, 0.8, nil
		} else {
			return 1.1, 0.6, 0.8, nil
		}

	case "5-4-1":
		if totalDefendersQuality >= 500 {
			return 1, 0.5, 0.5, nil
		} else {
			return 0.9, 0.5, 0.6, nil
		}

	case "5-3-2":
		if totalForwardersQuality >= 510 {
			return 0.7, 1.1, 0.8, nil
		} else {
			return 0.7, 1, 0.9, nil
		}

	case "3-4-3":
		if totalDefendersQuality >= 526 {
			return 1.2, 1.3, 1.3, nil
		} else {
			return 1, 1.3, 1.3, nil
		}

	case "3-5-2":
		if totalMidfieldersQuality >= 521 {
			return 1.2, 1.1, 1.1, nil
		} else {
			return 0.9, 1.1, 1.1, nil
		}

	default:
		return 0, 0, 0, errors.New("formación desconocida")
	}
}

func (a *AppService) CalculatePossessionChancesByPlayingStyle(lineup []Lineup, playingStyle string) (teamPossession, teamChances, rivalChances float64, physique int, err error) {
	totalDefendersQuality, err := getTwoBestPlayers(lineup, "defender")
	totalMidfieldersQuality, err := getTwoBestPlayers(lineup, "midfielder")
	totalForwardersQuality, err := getTwoBestPlayers(lineup, "forwarder")

	switch playingStyle {

	case "possession":
		if totalMidfieldersQuality >= 550 {
			return 1.6, 0.7, 0.8, 55, nil
		}
		return 1.4, 0.7, 0.8, 50, nil
	case "counter_attack":
		if totalForwardersQuality >= 470 {
			return 0.7, 1.3, 0.9, -15, nil
		}
		return 0.7, 1.2, 0.9, -20, nil
	case "direct_play":
		if totalForwardersQuality >= 400 {
			return 0.5, 1.1, 0.8, 20, nil
		}
		return 0.5, 1, 0.8, 10, nil
	case "high_press":
		if totalMidfieldersQuality >= 440 {
			return 1.1, 1.4, 1.12, -190, nil
		}
		return 1.1, 1.35, 1.12, -220, nil
	case "low_block":
		if totalDefendersQuality >= 410 {
			return 0.8, 0.4, 0.5, 130, nil
		}
		return 0.8, 0.3, 0.5, 110, nil
	default:
		return 0, 0, 0, 0, errors.New("playingStyle desconocido")
	}
}

func (a *AppService) CalculatePossessionChancesByGameTempo(gameTempo string) (teamPossession, teamChances, rivalChances float64, physique int, err error) {
	switch gameTempo {
	case "fast_tempo":
		return 0.8, 1.2, 1.1, -150, nil
	case "balanced_tempo":
		return 1, 1, 1, 10, nil
	case "slow_tempo":
		return 1.1, 0.6, 0.7, 250, nil

	default:
		return 0, 0, 0, 0, errors.New("gameTempo desconocido")
	}
}

func (a *AppService) CalculatePossessionChancesByPassingStyle(passingStyle string) (teamPossession, rivalChances float64, err error) {
	switch passingStyle {
	case "short":
		return 1.1, 1, nil
	case "long":
		return 0.8, 0.9, nil

	default:
		return 0, 0, errors.New("gameTempo desconocido")
	}
}

func (a *AppService) CalculateRivalChancesByDefensivePositioning(lineup []Lineup, defensivePositioning string) (rivalChances float64, physique int, err error) {

	var totalMentalityOfDefenders, totalPhysiqueOfDefenders int

	for _, player := range lineup {
		if player.Position == "defender" {
			totalMentalityOfDefenders += player.Mental
			totalPhysiqueOfDefenders += player.Physique
		}
	}

	switch defensivePositioning {
	case "zonal_marking":
		if totalMentalityOfDefenders >= 370 {
			return 0.7, 65, nil
		}
		if totalMentalityOfDefenders >= 290 {
			return 0.9, 40, nil
		}
		if totalMentalityOfDefenders >= 200 {
			return 1, 2, nil
		}

		return 1.45, -20, nil

	case "man_marking":
		if totalMentalityOfDefenders >= 340 {
			return 0.8, 15, nil
		}
		if totalMentalityOfDefenders >= 250 {
			return 0.9, 1, nil
		}
		if totalPhysiqueOfDefenders >= 190 {
			return 1, -40, nil
		}

		return 1.3, -120, nil

	default:
		return 0, 0, errors.New("Estilo de posicionamiento defensivo desconocido")
	}
}

func (a *AppService) CalculatePossessionByBuildUpPlay(buildUpPlay string) (possession float64, err error) {
	lineup, err := a.lineupRepo.GetLineup()
	if err != nil {
		return 0, errors.New("Error al obtener la alineación")
	}

	var totalTechniqueOfGoalkeeper, totalMentalityOfGoalkeeper, totalTechniqueOfDefenders, totalMentalOfDefenders int
	var defenderCount int

	for _, player := range lineup {
		if player.Position == "goalkeeper" {
			totalTechniqueOfGoalkeeper += player.Technique
			totalMentalityOfGoalkeeper += player.Mental
		} else if player.Position == "defender" {
			totalTechniqueOfDefenders += player.Technique
			totalMentalOfDefenders += player.Mental
			defenderCount++

		}
	}

	if defenderCount == 0 {
		return 0, errors.New("No hay defensores en la alineación")
	}

	totalQualityOfGoalkeeper := totalTechniqueOfGoalkeeper + totalMentalityOfGoalkeeper
	averageTotalQualityOfDefenders := (totalTechniqueOfDefenders + totalMentalOfDefenders) / defenderCount

	switch buildUpPlay {
	case "play_from_back":
		if totalTechniqueOfGoalkeeper >= 84 && totalMentalityOfGoalkeeper >= 84 && averageTotalQualityOfDefenders >= 79 {
			return 1.3, nil
		}
		if totalTechniqueOfGoalkeeper >= 82 && totalMentalityOfGoalkeeper >= 82 || averageTotalQualityOfDefenders >= 70 && totalQualityOfGoalkeeper >= 150 {
			return 1.23, nil
		}
		if totalQualityOfGoalkeeper >= 139 || averageTotalQualityOfDefenders >= 72 {
			return 1.10, nil
		}
		if totalTechniqueOfGoalkeeper >= 66 || totalMentalityOfGoalkeeper >= 66 {
			return 1.07, nil
		}

		return 0.63, nil

	case "long_clearance":

		if averageTotalQualityOfDefenders >= 86 {
			return 1.1, nil
		}
		if averageTotalQualityOfDefenders >= 74 {
			return 1.02, nil
		}

		return 0.9, nil

	default:
		return 0, errors.New("Estilo de buildUpPlay desconocido")
	}
}

func (a *AppService) CalculateRivalChancesByAttackFocus(lineup []Lineup, attackFocus string) (chances float64, err error) {

	var totalTechniqueOfMidfield, totalPhysiqueOfMidfild int
	var forwardCount, midfieldersCount int

	for _, player := range lineup {
		if player.Position == "midfielder" {
			totalTechniqueOfMidfield += player.Technique
			totalPhysiqueOfMidfild += player.Physique
			midfieldersCount++
		} else if player.Position == "forward" {
			forwardCount++

		}
	}

	if forwardCount == 0 {
		return 0, errors.New("No hay delanteros en la alineación")
	}

	totalQualityOfMidfield := totalTechniqueOfMidfield + totalPhysiqueOfMidfild
	averageTotalQualityOfMidfield := totalQualityOfMidfield / midfieldersCount

	switch attackFocus {
	case "wide_play":
		if averageTotalQualityOfMidfield >= 84 && forwardCount >= 2 {
			return 1.28, nil
		}
		if averageTotalQualityOfMidfield >= 82 {
			return 1.22, nil
		}
		if totalQualityOfMidfield >= 245 || forwardCount >= 2 {
			return 1.09, nil
		}
		if totalQualityOfMidfield >= 215 {
			return 1.06, nil
		}

		return 0.83, nil

	case "central_play":
		if averageTotalQualityOfMidfield >= 79 && midfieldersCount >= 4 {
			return 1.21, nil
		}
		if averageTotalQualityOfMidfield >= 76 {
			return 1.14, nil
		}
		if midfieldersCount >= 4 {
			return 1.09, nil
		}

		return 0.91, nil

	default:
		return 0, errors.New("Estilo de AttackFocus desconocido")
	}
}

func (a *AppService) CalculateRivalChancesByKeyPlayerUsage(lineup []Lineup, keyPlayerUsage string) (possession, chances float64, err error) {

	var keyPlayer Lineup

	for _, player := range lineup {
		if player.Technique+player.Mental+player.Physique > keyPlayer.Technique+keyPlayer.Mental+keyPlayer.Physique {
			keyPlayer = player
		}
	}

	totalQualityOfKeyPlayer := keyPlayer.Technique + keyPlayer.Mental + keyPlayer.Physique

	switch keyPlayerUsage {
	case "reference_player":
		if totalQualityOfKeyPlayer >= 278 {
			return 0.98, 1.9, nil
		}
		if totalQualityOfKeyPlayer >= 271 {
			return 0.94, 1.64, nil
		}
		if totalQualityOfKeyPlayer >= 254 {
			return 1, 1.52, nil
		}
		if totalQualityOfKeyPlayer >= 216 {
			return 1, 1.26, nil
		}
		if totalQualityOfKeyPlayer >= 204 {
			return 0.98, 1.1, nil
		}

		return 1.1, 0.67, nil

	case "free_role_player":

		return 1.3, 0.98, nil

	default:
		return 0, 0, errors.New("Estilo de KeyPlayerUsage desconocido")
	}
}

func getTwoBestPlayers(players []Lineup, position string) (int, error) {
	var foundPlayers int
	bestPlayers := make([]Lineup, 2)

	for _, player := range players {
		if player.Position == position {
			if foundPlayers < 2 {
				foundPlayers++
				if foundPlayers == 1 || player.Technique+player.Mental+player.Physique > bestPlayers[0].Technique+bestPlayers[0].Mental+bestPlayers[0].Physique {
					bestPlayers[1] = bestPlayers[0]
					bestPlayers[0] = player
				} else if foundPlayers == 2 || player.Technique+player.Mental+player.Physique > bestPlayers[1].Technique+bestPlayers[1].Mental+bestPlayers[1].Physique {
					bestPlayers[1] = player
				}
			}
		}
	}

	if foundPlayers < 2 {
		return 0, errors.New("no hay suficientes jugadores en la alineación")
	}

	totalPlayersQuality := bestPlayers[0].Technique + bestPlayers[0].Mental + bestPlayers[0].Physique +
		bestPlayers[1].Technique + bestPlayers[1].Mental + bestPlayers[1].Physique

	return totalPlayersQuality, nil
}
