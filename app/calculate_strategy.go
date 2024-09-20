package app

import "errors"

func (a *AppService) CalculatePossessionChancesByFormation(formation string) (teamPossession, teamChances, rivalChances float64, err error) {
	lineup, err := a.lineupRepo.GetLineup()
	if err != nil {
		return 0, 0, 0, errors.New("Error al obtener la alineación")
	}

	totalDefendersQuality, err := getTwoBestPlayers(lineup, "defense")
	totalMidfieldersQuality, err := getTwoBestPlayers(lineup, "midfield")
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

func (a *AppService) CalculatePossessionChancesByPlayingStyle(playingStyle string) (teamPossession, teamChances, rivalChances float64, physique int, err error) {
	switch playingStyle {
	case "possession":
		return 1.4, 0.7, 0.8, 50, nil
	case "counter_attack":
		return 0.7, 1.2, 0.9, -20, nil
	case "direct_play":
		return 0.5, 1, 0.8, 10, nil
	case "high_press":
		return 1.1, 1.35, 1.12, -220, nil
	case "low_block":
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

func (a *AppService) CalculateRivalChancesByDefensivePositioning(defensivePositioning string) (rivalChances float64, physique int, err error) {

	lineup, err := a.lineupRepo.GetLineup()
	if err != nil {
		return 0, 0, errors.New("Error al obtener la alineación")
	}

	var totalMentalityOfDefenses, totalPhysiqueOfDefenses int

	for _, player := range lineup {
		if player.Position == "defense" {
			totalMentalityOfDefenses += player.Mental
			totalPhysiqueOfDefenses += player.Physique
		}
	}

	switch defensivePositioning {
	case "zonal_marking":
		if totalMentalityOfDefenses >= 370 {
			return 0.7, 65, nil
		}
		if totalMentalityOfDefenses >= 290 {
			return 0.9, 40, nil
		}
		if totalMentalityOfDefenses >= 200 {
			return 1, 2, nil
		}

		return 1.45, -20, nil

	case "man_marking":
		if totalMentalityOfDefenses >= 340 {
			return 0.8, 15, nil
		}
		if totalMentalityOfDefenses >= 250 {
			return 0.9, 1, nil
		}
		if totalPhysiqueOfDefenses >= 190 {
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

	var totalTechniqueOfGoalkeeper, totalMentalityOfGoalkeeper, totalTechniqueOfDefenses, totalMentalOfDefenses int
	var defenseCount int

	for _, player := range lineup {
		if player.Position == "goalkeeper" {
			totalTechniqueOfGoalkeeper += player.Technique
			totalMentalityOfGoalkeeper += player.Mental
		} else if player.Position == "defense" {
			totalTechniqueOfDefenses += player.Technique
			totalMentalOfDefenses += player.Mental
			defenseCount++

		}
	}

	if defenseCount == 0 {
		return 0, errors.New("No hay defensores en la alineación")
	}

	totalQualityOfGoalkeeper := totalTechniqueOfGoalkeeper + totalMentalityOfGoalkeeper
	averageTotalQualityOfDefenses := (totalTechniqueOfDefenses + totalMentalOfDefenses) / defenseCount

	switch buildUpPlay {
	case "play_from_back":
		if totalTechniqueOfGoalkeeper >= 84 && totalMentalityOfGoalkeeper >= 84 && averageTotalQualityOfDefenses >= 79 {
			return 1.3, nil
		}
		if totalTechniqueOfGoalkeeper >= 82 && totalMentalityOfGoalkeeper >= 82 || averageTotalQualityOfDefenses >= 70 && totalQualityOfGoalkeeper >= 150 {
			return 1.23, nil
		}
		if totalQualityOfGoalkeeper >= 139 || averageTotalQualityOfDefenses >= 72 {
			return 1.10, nil
		}
		if totalTechniqueOfGoalkeeper >= 66 || totalMentalityOfGoalkeeper >= 66 {
			return 1.07, nil
		}

		return 0.63, nil

	case "long_clearance":

		if averageTotalQualityOfDefenses >= 86 {
			return 1.1, nil
		}
		if averageTotalQualityOfDefenses >= 74 {
			return 1.02, nil
		}

		return 0.9, nil

	default:
		return 0, errors.New("Estilo de buildUpPlay desconocido")
	}
}

func (a *AppService) CalculateRivalChancesByAttackFocus(attackFocus string) (chances float64, err error) {
	lineup, err := a.lineupRepo.GetLineup()
	if err != nil {
		return 0, errors.New("Error al obtener la alineación")
	}

	var totalTechniqueOfMidfield, totalPhysiqueOfMidfild int
	var forwardCount, midfieldCount int

	for _, player := range lineup {
		if player.Position == "midfield" {
			totalTechniqueOfMidfield += player.Technique
			totalPhysiqueOfMidfild += player.Physique
			midfieldCount++
		} else if player.Position == "forward" {
			forwardCount++

		}
	}

	if forwardCount == 0 {
		return 0, errors.New("No hay delanteros en la alineación")
	}

	totalQualityOfMidfield := totalTechniqueOfMidfield + totalPhysiqueOfMidfild
	averageTotalQualityOfMidfield := totalQualityOfMidfield / midfieldCount

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
		if averageTotalQualityOfMidfield >= 79 && midfieldCount >= 4 {
			return 1.21, nil
		}
		if averageTotalQualityOfMidfield >= 76 {
			return 1.14, nil
		}
		if midfieldCount >= 4 {
			return 1.09, nil
		}

		return 0.91, nil

	default:
		return 0, errors.New("Estilo de AttackFocus desconocido")
	}
}

func (a *AppService) CalculateRivalChancesByKeyPlayerUsage(keyPlayerUsage string) (possession, chances float64, err error) {
	lineup, err := a.lineupRepo.GetLineup()
	if err != nil {
		return 0, 0, errors.New("Error al obtener la alineación")
	}

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
