package app

import (
	"log"
)

func (a *AppService) ProcessMatchPlay() (Match, error) {

	lineup, err := a.lineupRepo.GetLineup()
	if err != nil {
		log.Println("error al obtener lineup", err)
		return Match{}, err
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

	numberOfLineupEvents, numberOfRivalEvents, err := DistributeMatchEvents(lineup, rivalLineup, numberOfMatchEvents, homeOrAwayString)
	if err != nil {
		log.Println("error al distribuir numberOfMatchEvents", err)
	}

	a.GenerateEvents(lineup, rivalLineup, numberOfLineupEvents, numberOfRivalEvents)

	Match := Match{}

	return Match, nil
}
