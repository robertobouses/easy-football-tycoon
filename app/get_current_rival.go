package app

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

func (a *AppService) GetCurrentRival() (Rival, string, int, error) {
	if a.currentRivals == nil {
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
	if totalRivals == 0 {
		return Rival{}, "", 0, errors.New("no hay rivales disponibles")
	}

	if a.callCounterRival <= 0 || a.callCounterRival > 2*totalRivals {
		log.Printf("Índice de rival fuera de rango: callCounterRival=%d, totalRivals=%d", a.callCounterRival, totalRivals)
		return Rival{}, "", 0, errors.New("índice de rival fuera de rango")
	}

	isSecondRound := a.callCounterRival > totalRivals
	index := (a.callCounterRival - 1) % totalRivals

	homeOrAway := "home"
	if isSecondRound {
		homeOrAway = "away"
	} else if rand.Intn(2) == 0 {
		homeOrAway = "away"
	}

	currentRival := (*a.currentRivals)[index]
	a.callCounterRival++

	if a.callCounterRival > 2*totalRivals {
		a.callCounterRival = 1
	}

	return currentRival, homeOrAway, a.callCounterRival - 1, nil
}
