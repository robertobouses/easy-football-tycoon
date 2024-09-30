package app

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func (a *AppService) ProcessMatch(calendaryId int) (Match, error) {
	lineup, err := a.lineupRepo.GetLineup()
	if err != nil {
		return Match{}, err
	}

	if len(lineup) == 0 {
		return Match{}, errors.New("no se encontraron jugadores en la alineación")
	}
	if len(lineup) < 11 {
		return Match{}, ErrLineupIncompleted
	}

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

	totalTechnique := lineup[0].TotalTechnique
	totalMental := lineup[0].TotalMental
	totalPhysique := lineup[0].TotalPhysique

	strategy, err := a.GetStrategy()
	if err != nil {
		log.Println("Error al obtener la estrategia:", err)
		return Match{}, fmt.Errorf("error al obtener la estrategia: %w", err)
	}
	resultOfStrategy, err := a.ResultOfStrategy(lineup, strategy.Formation, strategy.PlayingStyle, strategy.GameTempo, strategy.PassingStyle, strategy.DefensivePositioning, strategy.BuildUpPlay, strategy.AttackFocus, strategy.KeyPlayerUsage)
	if err != nil {
		log.Println("Error al calcular el resultado de la estrategia:", err)
		return Match{}, fmt.Errorf("error al calcular el resultado de la estrategia: %w", err)
	}

	totalPhysique = totalPhysique + int(resultOfStrategy["teamPhysique"])

	lineupTotalQuality, rivalTotalQuality, allQuality, err := CalculateTotalQuality(totalTechnique, totalMental, totalPhysique, rival.Technique, rival.Mental, rival.Physique)
	if err != nil {
		return Match{}, err
	}

	lineupPercentagePossession, rivalPercentagePossession, err := a.CalculateBallPossession(totalTechnique, rival.Technique, lineupTotalQuality, rivalTotalQuality, allQuality, homeOrAwayString, resultOfStrategy["teamPossession"])
	if err != nil {
		return Match{}, err
	}

	lineupScoreChances, rivalScoreChances, err := a.CalculateScoringChances(lineupTotalQuality, rivalTotalQuality, totalMental, rival.Mental, homeOrAwayString, resultOfStrategy["teamChances"], resultOfStrategy["rivalChances"])
	if err != nil {
		return Match{}, err
	}

	lineupGoals, rivalGoals, err := CalculateGoals(totalPhysique, rival.Physique, lineupPercentagePossession, rivalPercentagePossession, lineupScoreChances, rivalScoreChances)
	if err != nil {
		return Match{}, err
	}

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
		Result: result,
	}

	a.SetCurrentMatch(&match)

	err = a.matchRepo.PostMatch(match)
	if err != nil {
		log.Println("Error en el procesamiento del partido:", err)
		return Match{}, ErrProcessingMatch
	}

	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("Error al obtener analíticas en ProcessMatch:", err)
		return Match{}, err
	}

	matches, err := a.matchRepo.GetMatches()
	if err != nil {
		log.Println("Error al obtener partidos en ProcessMatch:", err)
		return Match{}, err
	}

	var previousMatchScoreChances int
	var previousMatchResult string

	if len(matches) > 0 {
		lastMatch := matches[len(matches)-1]
		previousMatchScoreChances = lastMatch.Team.ScoringChances
		previousMatchResult = lastMatch.Result
	} else {
		log.Println("No hay partidos previos disponibles")

		previousMatchScoreChances = 0
		previousMatchResult = "DRAW"
	}

	log.Println("stadiumCapacity en ProcessMatch:", analytics.StadiumCapacity)
	log.Println("currentTrust en ProcessMatch:", analytics.Trust)

	ticketSales, err := a.CalculateTicketSales(analytics.StadiumCapacity, lineupTotalQuality, rivalTotalQuality, analytics.Trust, previousMatchScoreChances, previousMatchResult, homeOrAwayString)
	if err != nil {
		log.Println("Error en el cálculo de tickets del estadio:", err)
		return Match{}, err
	}

	initialBalance, err := a.bankRepo.GetBalance()
	if err != nil {
		return Match{}, ErrBalanceNotFound
	}

	newBalance := initialBalance + ticketSales
	err = a.bankRepo.PostTransactions(ticketSales, newBalance, "Ticket Sales", "Ticket Sales")
	if err != nil {
		log.Println("Error al registrar la transacción:", err)
		return Match{}, err
	}

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

func (a *AppService) CalculateBallPossession(lineupTotalTechnique, rivalTotalTechnique, lineupTotalQuality, rivalTotalQuality, allQuality int, homeOrAway string, lineupPossessionResultOfStrategy float64) (int, int, error) {
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
	log.Println("posesión equipo antes de strategia", percentageLineupQuality)

	percentageLineupQuality = percentageLineupQuality * lineupPossessionResultOfStrategy

	log.Println("posesión equipo después de strategia", percentageLineupQuality)

	rand.Seed(time.Now().UnixNano())
	randomFactor := 0.8 + rand.Float64()*(1.2-0.8)
	log.Println("el randomFactor es", randomFactor)
	percentageLineupQualityWithRandomFactor := percentageLineupQuality * randomFactor
	log.Println("posesión equipo después de randomFactor", percentageLineupQualityWithRandomFactor)

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
	log.Println("posesión equipo después de Home or Away", percentageLineupQualityWithRandomFactor)

	if percentageLineupQualityWithRandomFactor > 83 {
		percentageLineupQualityWithRandomFactor = 83
	} else if percentageLineupQualityWithRandomFactor < 17 {
		percentageLineupQualityWithRandomFactor = 17
	}

	percentageLineup := int(percentageLineupQualityWithRandomFactor)
	percentageRival := 100 - percentageLineup

	return percentageLineup, percentageRival, nil
}

func (a *AppService) CalculateScoringChances(lineupTotalQuality, rivalTotalQuality, lineupTotalMental, rivalTotalMental int, homeOrAway string, lineupScoringChancesResultOfStrategy, rivalScoringChancesResultOfStrategy float64) (int, int, error) {
	var lineupScoringChances, rivalScoringChances float64

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
	switch {
	case float64(lineupTotalMental) >= float64(rivalTotalMental)*1.5:
		lineupScoringChances = lineupScoringChances * 1.25
		rivalScoringChances = rivalScoringChances * 0.3
	case float64(lineupTotalMental) >= float64(rivalTotalMental)*1.4:
		lineupScoringChances = lineupScoringChances * 1.1
		rivalScoringChances = rivalScoringChances * 0.4
	case float64(lineupTotalMental) >= float64(rivalTotalMental)*1.3:
		lineupScoringChances = lineupScoringChances * 1.05
		rivalScoringChances = rivalScoringChances * 0.5
	case float64(lineupTotalMental) >= float64(rivalTotalMental)*1.2:
		lineupScoringChances = lineupScoringChances * 1.04
		rivalScoringChances = rivalScoringChances * 0.6
	case float64(lineupTotalMental) >= float64(rivalTotalMental)*1.1:
		lineupScoringChances = lineupScoringChances * 1.02
		rivalScoringChances = rivalScoringChances * 0.7
	}

	switch {
	case float64(lineupTotalMental) <= float64(rivalTotalMental)*1.5:
		rivalScoringChances = rivalScoringChances * 1.25
		lineupScoringChances = lineupScoringChances * 0.3
	case float64(lineupTotalMental) <= float64(rivalTotalMental)*1.4:
		rivalScoringChances = rivalScoringChances * 1.1
		lineupScoringChances = lineupScoringChances * 0.4
	case float64(lineupTotalMental) <= float64(rivalTotalMental)*1.3:
		rivalScoringChances = rivalScoringChances * 1.05
		lineupScoringChances = lineupScoringChances * 0.5
	case float64(lineupTotalMental) <= float64(rivalTotalMental)*1.2:
		rivalScoringChances = rivalScoringChances * 1.04
		lineupScoringChances = lineupScoringChances * 0.8
	case float64(lineupTotalMental) <= float64(rivalTotalMental)*1.1:
		rivalScoringChances = rivalScoringChances * 1.02
		lineupScoringChances = lineupScoringChances * 0.7
	}
	log.Println("lineupScoringChances antes de strategy", lineupScoringChances)
	log.Println("rivalScoringChances antes de strategy", rivalScoringChances)

	lineupScoringChances = lineupScoringChances * lineupScoringChancesResultOfStrategy
	rivalScoringChances = rivalScoringChances * rivalScoringChancesResultOfStrategy

	log.Println("lineupScoringChances después de strategy", lineupScoringChances)
	log.Println("rivalScoringChances después de strategy", rivalScoringChances)

	rand.Seed(time.Now().UnixNano())
	randomFactor := 0.84 + rand.Float64()*(2.3-0.84)
	log.Println("el randomFactor es", randomFactor)

	lineupScoringChances = lineupScoringChances * randomFactor
	rivalScoringChances = rivalScoringChances * randomFactor

	log.Println("lineupScoringChances después de randomFactor", lineupScoringChances)
	log.Println("rivalScoringChances después de randomFactor", rivalScoringChances)

	if homeOrAway == "home" {
		lineupScoringChances = lineupScoringChances * 1.5
	} else if homeOrAway == "away" {
		rivalScoringChances = rivalScoringChances * 1.5
	}

	log.Println("lineupScoringChances después de Home or Away", lineupScoringChances)
	log.Println("rivalScoringChances después de Home or Away", rivalScoringChances)

	if lineupScoringChances <= 2 {
		randomFactor = 0 + rand.Float64()*(3-0)
		log.Println("El segundo randomFactor, utilizado si las ocasiones son nulas o escasas del Equipo, es:", randomFactor)
		lineupScoringChances = 1.00
		lineupScoringChances = lineupScoringChances * randomFactor
	}

	if rivalScoringChances <= 2 {
		randomFactor = 0 + rand.Float64()*(3-0)
		log.Println("El segundo randomFactor, utilizado si las ocasiones son nulas o escasas del Rival, es:", randomFactor)
		rivalScoringChances = 1.00
		rivalScoringChances = rivalScoringChances * randomFactor
	}

	if lineupScoringChances > 29 {
		lineupScoringChances = 29
	}
	if rivalScoringChances > 29 {
		rivalScoringChances = 29
	}
	if lineupScoringChances > 19 {
		randomFactor := 0.81 + rand.Float64()*(1.01-0.81)
		log.Println("El tercer randomFactor, utilizado si las ocasiones son muy altas del Equipo, es:", randomFactor)

		lineupScoringChances = lineupScoringChances * randomFactor
		log.Println("Ocasiones del equipo después de aplicar el factor:", lineupScoringChances)
	}

	if rivalScoringChances > 19 {
		randomFactor := 0.81 + rand.Float64()*(1.01-0.81)

		log.Println("El tercer randomFactor, utilizado si las ocasiones  son muy altas del Rival, es:", randomFactor)

		rivalScoringChances = rivalScoringChances * randomFactor
		log.Println("Ocasiones del rival después de aplicar el factor:", rivalScoringChances)
	}

	return int(lineupScoringChances), int(rivalScoringChances), nil

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

			probability := float64(possession)*0.5 + float64(physique)*0.009

			if chance < probability {
				goals++
			}
		}

		return goals
	}

	randomFactor := 0.71 + rand.Float64()*(1.01-0.71)

	goalsLineup := calculateGoals(lineupScoringChances, lineupPossession, lineupTotalPhysique)
	if goalsLineup > 12 {
		goalsLineup = 12
	}
	if goalsLineup >= 8 {
		goalsLineup = int(float64(goalsLineup) * randomFactor)
	}
	if goalsLineup > 5 {
		restarGol := rand.Intn(4) < 3
		if restarGol {
			goalsLineup--
		}
	}

	goalsRival := calculateGoals(rivalScoringChances, rivalPossession, rivalTotalPhysique)
	if goalsRival > 12 {
		goalsRival = 12
	}
	if goalsRival >= 8 {
		goalsRival = int(float64(goalsRival) * randomFactor)
	}
	if goalsRival > 5 {
		restarGol := rand.Intn(4) < 3
		if restarGol {
			goalsRival--
		}
	}

	if goalsLineup > 2 && goalsRival > 2 && (goalsLineup+goalsRival > 8) {
		goalsLineup--
		goalsRival--
	}

	if goalsLineup > 2 && goalsRival > 2 && (goalsLineup+goalsRival > 6) {
		restarGol := rand.Intn(2) == 0
		if restarGol {
			goalsLineup--
			goalsRival--
		}
	}

	if goalsLineup > 4 {
		restarGol := rand.Intn(3) == 0
		if restarGol {
			goalsLineup--

		}
	}
	if goalsRival > 4 {
		restarGol := rand.Intn(3) == 0
		if restarGol {
			goalsRival--

		}

	}

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
