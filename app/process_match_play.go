package app

import (
	"log"
	"math/rand"
	"sort"
)

func (a *AppService) ProcessMatchPlay() (Match, error) {

	lineup, err := a.lineupRepo.GetLineup()
	if err != nil {
		log.Println("error al obtener lineup", err)
	}

	rival, homeOrAwayString, matchDayNumber, err := a.GetCurrentRival()
	if err != nil {
		log.Println("Error al obtener el rival:", err)
		return Match{}, err
	}
	log.Printf("Rival seleccionado: %s\n", rival.RivalName)
	log.Printf("Día del partido: %d\n", matchDayNumber)

	rivalLineup := a.SimulateRivalLineup(rival)

	numberOfMatchEvents, err := a.CalculateNumberOfMatchEvents()
	if err != nil {
		log.Println("error al obtener numberOfMatchEvents", err)
	}

	matchClockEvents := a.SimulateMatchClock(numberOfMatchEvents)

	lineupEvents, rivalEvents, err := DistributeMatchEvents(lineup, rivalLineup, numberOfMatchEvents, homeOrAwayString)
	if err != nil {
		log.Println("error al distribuir numberOfMatchEvents", err)
	}

	events := append(lineupEvents, rivalEvents...)
	for _, event := range events {
		matchClockEvents = append(matchClockEvents, event)
	}

	sort.Slice(matchClockEvents, func(i, j int) bool {
		return matchClockEvents[i].Minute < matchClockEvents[j].Minute
	})

	var passer, receiver, shooter, goalkeeper, defender Player
	var lastEvent string

	for _, matchEvent := range matchClockEvents {
		log.Printf("Evento: %s en el minuto %d", matchEvent.Description, matchEvent.Minute)

		switch matchEvent.Description {
		case "KeyPass":
			if _, err := a.KeyPass(passer, receiver); err != nil {
				log.Printf("Error en KeyPass: %v", err)
			}
			lastEvent = "KeyPass"

		case "Shot":
			if lastEvent == "KeyPass" {
				continue
			}
			if _, err := a.Shot(shooter, goalkeeper, defender, &passer); err != nil {
				log.Printf("Error en Shot: %v", err)
			}
			lastEvent = "Shot"

		case "LongShot":
			if _, err := a.LongShot(shooter, goalkeeper); err != nil {
				log.Printf("Error en LongShot: %v", err)
			}
			lastEvent = "LongShot"

		case "PenaltyKick":
			if lastEvent == "KeyPass" {
				if _, err := a.PenaltyKick(shooter, goalkeeper); err != nil {
					log.Printf("Error en PenaltyKick: %v", err)
				}
				lastEvent = "PenaltyKick"
			}
		}

		if lastEvent == "KeyPass" {
			if rand.Float32() < 0.8 {
				if _, err := a.Shot(shooter, goalkeeper, defender, &passer); err != nil {
					log.Printf("Error en Shot después de KeyPass: %v", err)
				}
			} else {
				if _, err := a.PenaltyKick(shooter, goalkeeper); err != nil {
					log.Printf("Error en PenaltyKick después de KeyPass: %v", err)
				}
			}
		}
	}

	// Devuelve el resultado del partido (puedes ajustar esto según tus necesidades)
	return Match{}, nil
}
