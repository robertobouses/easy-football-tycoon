package app

import (
	"log"
	"math/rand"
	"time"
)

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
