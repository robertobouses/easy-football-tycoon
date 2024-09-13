package app

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func (a *AppService) ProcessMatch(calendaryId int) (Match, error) {
	if a.currentRivals == nil || len(*a.currentRivals) == 0 {
		return Match{}, errors.New("currentRivals no está inicializado o está vacío")
	}

	rival, homeOrAwayString, matchDayNumber, err := a.GetCurrentRival()
	if err != nil {
		return Match{}, err
	}
	log.Println(matchDayNumber)

	var homeOrAway HomeOrAway
	switch homeOrAwayString {
	case "home":
		homeOrAway = Home
	case "away":
		homeOrAway = Away
	default:
		return Match{}, errors.New("valor inválido para homeOrAway")
	}

	lineup, err := a.GetLineup()
	if err != nil {
		return Match{}, err
	}

	if len(lineup) == 0 {
		return Match{}, errors.New("no se encontraron jugadores en la alineación")
	}

	totalTechnique := lineup[0].TotalTechnique
	totalMental := lineup[0].TotalMental
	totalPhysique := lineup[0].TotalPhysique

	lineupTotalQuality, rivalTotalQuality, allQuality, err := CalculateTotalQuality(totalTechnique, totalMental, totalPhysique, rival.Technique, rival.Mental, rival.Physique)
	if err != nil {
		return Match{}, err
	}

	lineupPercentagePossession, rivalPercentagePossession, err := CalculateBallPossession(totalTechnique, rival.Technique, lineupTotalQuality, rivalTotalQuality, allQuality, homeOrAwayString)
	if err != nil {
		return Match{}, err
	}

	lineupScoreChances, rivalScoreChances, err := CalculateScoringChances(lineupTotalQuality, rivalTotalQuality, totalMental, rival.Mental, homeOrAwayString)
	if err != nil {
		return Match{}, err
	}

	lineupGoals, rivalGoals, err := CalculateGoals(totalPhysique, rival.Physique, lineupPercentagePossession, rivalPercentagePossession, lineupScoreChances, rivalScoreChances)
	if err != nil {
		return Match{}, err
	}
	match := Match{
		MatchId:        uuid.New(),
		CalendaryId:    calendaryId,
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
	}

	a.SetCurrentMatch(&match)

	return match, nil
}

func CalculateTotalQuality(lineupTotalTechnique, lineupTotalMental, lineupTotalPhysique, rivalTotalTechnique, rivalTotalMental, rivalTotalPhysique int) (int, int, int, error) {

	lineupTotalQuality := lineupTotalTechnique + lineupTotalMental + lineupTotalPhysique
	rivalTotalQuality := rivalTotalTechnique + rivalTotalMental + rivalTotalPhysique
	allQuality := lineupTotalQuality + rivalTotalQuality

	if allQuality == 0 {
		return 0, 0, 0, ErrTeamsQualityNil
	}

	return lineupTotalQuality, rivalTotalQuality, allQuality, nil

}

func CalculateBallPossession(lineupTotalTechnique, rivalTotalTechnique, lineupTotalQuality, rivalTotalQuality, allQuality int, homeOrAway string) (int, int, error) {
	percentageLineupQuality := (float64(lineupTotalQuality) / float64(allQuality)) * 100

	switch {
	case float64(lineupTotalTechnique) >= float64(rivalTotalTechnique)*1.5:
		percentageLineupQuality *= 1.25
	case float64(lineupTotalTechnique) >= float64(rivalTotalTechnique)*1.4:
		percentageLineupQuality *= 1.2
	case float64(lineupTotalTechnique) >= float64(rivalTotalTechnique)*1.3:
		percentageLineupQuality *= 1.15
	case float64(lineupTotalTechnique) >= float64(rivalTotalTechnique)*1.2:
		percentageLineupQuality *= 1.1
	case float64(lineupTotalTechnique) >= float64(rivalTotalTechnique)*1.1:
		percentageLineupQuality *= 1.05
	}

	switch {
	case float64(lineupTotalTechnique) <= float64(rivalTotalTechnique)*1.5:
		percentageLineupQuality /= 1.25
	case float64(lineupTotalTechnique) <= float64(rivalTotalTechnique)*1.4:
		percentageLineupQuality /= 1.2
	case float64(lineupTotalTechnique) <= float64(rivalTotalTechnique)*1.3:
		percentageLineupQuality /= 1.15
	case float64(lineupTotalTechnique) <= float64(rivalTotalTechnique)*1.2:
		percentageLineupQuality /= 1.1
	case float64(lineupTotalTechnique) <= float64(rivalTotalTechnique)*1.1:
		percentageLineupQuality /= 1.05
	}

	rand.Seed(time.Now().UnixNano())
	randomFactor := 0.8 + rand.Float64()*(1.2-0.8)
	log.Println("el randomFactor es", randomFactor)
	percentageLineupQualityWithRandomFactor := percentageLineupQuality * randomFactor

	if homeOrAway == "home" {
		if percentageLineupQualityWithRandomFactor <= 45 {
			percentageLineupQualityWithRandomFactor *= 1.22
		} else if percentageLineupQualityWithRandomFactor <= 54 {
			percentageLineupQualityWithRandomFactor *= 1.15
		} else if percentageLineupQualityWithRandomFactor <= 67 {
			percentageLineupQualityWithRandomFactor *= 1.07
		}
	}

	if homeOrAway == "away" { //TODO NO DA EXACTO SIMÉTRICO
		if percentageLineupQualityWithRandomFactor >= 55 {
			percentageLineupQualityWithRandomFactor *= 0.9057
		} else if percentageLineupQualityWithRandomFactor >= 46 {
			percentageLineupQualityWithRandomFactor *= 0.9357
		} else if percentageLineupQualityWithRandomFactor >= 33 {
			percentageLineupQualityWithRandomFactor *= 0.97
		}
	}

	if percentageLineupQualityWithRandomFactor > 83 {
		percentageLineupQualityWithRandomFactor = 83
	} else if percentageLineupQualityWithRandomFactor < 17 {
		percentageLineupQualityWithRandomFactor = 17
	}

	percentageLineup := int(percentageLineupQualityWithRandomFactor)
	percentageRival := 100 - percentageLineup

	return percentageLineup, percentageRival, nil
}

func CalculateScoringChances(lineupTotalQuality, rivalTotalQuality, lineupTotalMental, rivalTotalMental int, homeOrAway string) (int, int, error) {
	lineupScoringChances := 0
	rivalScoringChances := 0
	switch {
	case lineupTotalQuality >= 3000:
		lineupScoringChances = 11
	case lineupTotalQuality >= 2750:
		lineupScoringChances = 10
	case lineupTotalQuality >= 2600:
		lineupScoringChances = 9
	case lineupTotalQuality >= 2400:
		lineupScoringChances = 8
	case lineupTotalQuality >= 2300:
		lineupScoringChances = 7
	case lineupTotalQuality >= 2200:
		lineupScoringChances = 6
	case lineupTotalQuality >= 2100:
		lineupScoringChances = 5
	case lineupTotalQuality >= 1900:
		lineupScoringChances = 4
	case lineupTotalQuality >= 1700:
		lineupScoringChances = 3
	case lineupTotalQuality >= 1400:
		lineupScoringChances = 2
	case lineupTotalQuality >= 0:
		lineupScoringChances = 1
	}

	switch {
	case rivalTotalQuality >= 3000:
		rivalScoringChances = 11
	case rivalTotalQuality >= 2750:
		rivalScoringChances = 10
	case rivalTotalQuality >= 2600:
		rivalScoringChances = 9
	case rivalTotalQuality >= 2400:
		rivalScoringChances = 8
	case rivalTotalQuality >= 2300:
		rivalScoringChances = 7
	case rivalTotalQuality >= 2200:
		rivalScoringChances = 6
	case rivalTotalQuality >= 2100:
		rivalScoringChances = 5
	case rivalTotalQuality >= 1900:
		rivalScoringChances = 4
	case lineupTotalQuality >= 1700:
		lineupScoringChances = 3
	case lineupTotalQuality >= 1400:
		lineupScoringChances = 2
	case lineupTotalQuality >= 0:
		lineupScoringChances = 1
	}
	//añadir aquí las modificaciones por ordenes, estrategia, formación
	switch {
	case float64(lineupTotalMental) >= float64(rivalTotalMental)*1.5:
		lineupScoringChances = int(float64(lineupScoringChances) * 1.25)
		rivalScoringChances = int(float64(rivalScoringChances) * 0.3)
	case float64(lineupTotalMental) >= float64(rivalTotalMental)*1.4:
		lineupScoringChances = int(float64(lineupScoringChances) * 1.1)
		rivalScoringChances = int(float64(rivalScoringChances) * 0.4)
	case float64(lineupTotalMental) >= float64(rivalTotalMental)*1.3:
		lineupScoringChances = int(float64(lineupScoringChances) * 1.05)
		rivalScoringChances = int(float64(rivalScoringChances) * 0.5)
	case float64(lineupTotalMental) >= float64(rivalTotalMental)*1.2:
		lineupScoringChances = int(float64(lineupScoringChances) * 1.04)
		rivalScoringChances = int(float64(rivalScoringChances) * 0.6)
	case float64(lineupTotalMental) >= float64(rivalTotalMental)*1.1:
		lineupScoringChances = int(float64(lineupScoringChances) * 1.02)
		rivalScoringChances = int(float64(rivalScoringChances) * 0.7)
	}

	switch {
	case float64(lineupTotalMental) <= float64(rivalTotalMental)*1.5:
		rivalScoringChances = int(float64(rivalScoringChances) * 1.25)
		lineupScoringChances = int(float64(lineupScoringChances) * 0.3)
	case float64(lineupTotalMental) <= float64(rivalTotalMental)*1.4:
		rivalScoringChances = int(float64(rivalScoringChances) * 1.1)
		lineupScoringChances = int(float64(lineupScoringChances) * 0.4)
	case float64(lineupTotalMental) <= float64(rivalTotalMental)*1.3:
		rivalScoringChances = int(float64(rivalScoringChances) * 1.05)
		lineupScoringChances = int(float64(lineupScoringChances) * 0.5)
	case float64(lineupTotalMental) <= float64(rivalTotalMental)*1.2:
		rivalScoringChances = int(float64(rivalScoringChances) * 1.04)
		lineupScoringChances = int(float64(lineupScoringChances) * 0.8)
	case float64(lineupTotalMental) <= float64(rivalTotalMental)*1.1:
		rivalScoringChances = int(float64(rivalScoringChances) * 1.02)
		lineupScoringChances = int(float64(lineupScoringChances) * 0.7)
	}

	rand.Seed(time.Now().UnixNano())
	randomFactor := 0.84 + rand.Float64()*(2.3-0.84)
	log.Println("el randomFactor es", randomFactor)

	lineupScoringChancesWithRandomFactor := int(float64(lineupScoringChances) * randomFactor)
	rivalScoringChancesWithRandomFactor := int(float64(rivalScoringChances) * randomFactor)

	if homeOrAway == "home" {
		lineupScoringChances = int(float64(lineupScoringChances) * 1.5)
	} else if homeOrAway == "away" {
		rivalScoringChances = int(float64(rivalScoringChances) * 1.5)
	}

	if lineupScoringChancesWithRandomFactor <= 2 {
		randomFactor = 0 + rand.Float64()*(3-0)
		log.Println("El segundo randomFactor es", randomFactor)
		lineupScoringChancesWithRandomFactor := 1
		lineupScoringChancesWithRandomFactor = int(float64(lineupScoringChancesWithRandomFactor) * randomFactor)
	}

	if rivalScoringChancesWithRandomFactor <= 2 {
		randomFactor = 0 + rand.Float64()*(3-0)
		log.Println("El segundo randomFactor es", randomFactor)
		rivalScoringChancesWithRandomFactor := 1
		rivalScoringChancesWithRandomFactor = int(float64(rivalScoringChancesWithRandomFactor) * randomFactor)
	}

	return lineupScoringChancesWithRandomFactor, rivalScoringChancesWithRandomFactor, nil

}

func CalculateGoals(lineupTotalPhysique, rivalTotalPhysique, lineupPossession, rivalPossession, lineupScoringChances, rivalScoringChances int) (int, int, error) {

	if lineupScoringChances == 0 && rivalScoringChances == 0 {
		return 0, 0, nil
	}

	rand.Seed(time.Now().UnixNano())

	calculateGoals := func(scoringChances, possession, physique int) int {
		goals := 0

		for i := 0; i < scoringChances; i++ {
			chance := rand.Float64() * 100

			probability := float64(possession)*0.5 + float64(physique)*0.015

			if chance < probability {
				goals++
			}
		}

		return goals
	}

	goalsLineup := calculateGoals(lineupScoringChances, lineupPossession, lineupTotalPhysique)
	goalsRival := calculateGoals(rivalScoringChances, rivalPossession, rivalTotalPhysique)

	return goalsLineup, goalsRival, nil
}

func (a *AppService) GetCurrentRival() (Rival, string, int, error) {
	if a.currentRivals == nil {
		return Rival{}, "", 0, errors.New("currentRivals es nil")
	}

	if len(*a.currentRivals) == 0 {
		return Rival{}, "", 0, errors.New("no hay rivales disponibles")
	}

	if a.callCounterRival <= 0 || a.callCounterRival > 2*len(*a.currentRivals) {
		return Rival{}, "", 0, errors.New("índice de rival fuera de rango")
	}

	if a.callCounterRival == 1 {
		rivals, err := a.rivalRepo.GetRival()
		if err != nil {
			return Rival{}, "", 0, err
		}

		if len(rivals) == 0 {
			return Rival{}, "", 0, errors.New("no se pudieron obtener rivales")
		}

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(rivals), func(i, j int) {
			rivals[i], rivals[j] = rivals[j], rivals[i]
		})

		a.currentRivals = &rivals
		a.callCounterRival = 1
	}

	totalRivals := len(*a.currentRivals)
	roundMatches := totalRivals
	isSecondRound := a.callCounterRival > roundMatches

	index := (a.callCounterRival - 1) % roundMatches

	homeOrAway := "home"
	if isSecondRound {
		homeOrAway = "away"
	}
	if !isSecondRound {
		if rand.Intn(2) == 0 {
			homeOrAway = "away"
		}
	}

	currentRival := (*a.currentRivals)[index]
	a.callCounterRival++

	if a.callCounterRival > 2*roundMatches {
		a.callCounterRival = 1
	}

	return currentRival, homeOrAway, a.callCounterRival - 1, nil
}

func (a *AppService) SetCurrentMatch(match *Match) {
	a.currentMatch = match
}

func (a *AppService) GetCurrentMatch() (*Match, error) {
	return a.currentMatch, nil
}
