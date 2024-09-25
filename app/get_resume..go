package app

import (
	"log"
)

func (a *AppService) GetResume() ([]Calendary, error) {
	lineup, err := a.lineupRepo.GetLineup()
	if err != nil {
		return []Calendary{}, err
	}

	if len(lineup) != 11 {
		log.Println("La alineación debe tener 11 jugadores. Deteniendo la ejecución.")
		return []Calendary{}, ErrLineupIncompleted
	}

	var goalkeeperCount, defenseCount, midfieldCount, forwardCount int

	//TODO COMPROBACIONES DE ALINEACION EN UNA FUNCIÓN FUERA DE GET RESUME

	for _, player := range lineup {
		switch player.Position {
		case "goalkeeper":
			goalkeeperCount++
		case "defense":
			defenseCount++
		case "midfield":
			midfieldCount++
		case "forward":
			forwardCount++
		default:
			log.Printf("Posición desconocida: %v", player.Position)
		}
	}

	strategy, err := a.strategyRepo.GetStrategy()
	if err != nil {
		return nil, err
	}
	switch strategy.Formation {
	case "4-4-2":

		if goalkeeperCount != 1 || defenseCount != 4 || midfieldCount != 4 || forwardCount != 2 {
			log.Println("La alineación no cumple con la formación 4-4-2.")
			return []Calendary{}, ErrInvalidFormation
		}
	case "4-3-3":

		if goalkeeperCount != 1 || defenseCount != 4 || midfieldCount != 3 || forwardCount != 3 {
			log.Println("La alineación no cumple con la formación 4-3-3.")
			return []Calendary{}, ErrInvalidFormation
		}

	case "4-5-1":

		if goalkeeperCount != 1 || defenseCount != 4 || midfieldCount != 5 || forwardCount != 1 {
			log.Println("La alineación no cumple con la formación 4-5-1.")
			return []Calendary{}, ErrInvalidFormation
		}
	case "5-3-2":

		if goalkeeperCount != 1 || defenseCount != 5 || midfieldCount != 3 || forwardCount != 2 {
			log.Println("La alineación no cumple con la formación 5-3-2.")
			return []Calendary{}, ErrInvalidFormation
		}
	case "3-4-3":

		if goalkeeperCount != 1 || defenseCount != 3 || midfieldCount != 4 || forwardCount != 3 {
			log.Println("La alineación no cumple con la formación 3-4-3.")
			return []Calendary{}, ErrInvalidFormation
		}
	case "3-5-2":

		if goalkeeperCount != 1 || defenseCount != 3 || midfieldCount != 5 || forwardCount != 2 {
			log.Println("La alineación no cumple con la formación 3-5-2.")
			return []Calendary{}, ErrInvalidFormation
		}
	default:
		log.Println("Formación desconocida:", strategy.Formation)
		return []Calendary{}, ErrUnknownFormation
	}
	a.callCounter++
	calendary, err := a.calendaryRepo.GetCalendary()
	if err != nil {
		log.Println("Error al extraer GetCalendary:", err)
		return []Calendary{}, err
	}

	if a.callCounter > len(calendary) {
		log.Println("No hay más fechas disponibles en el calendario")
		return []Calendary{}, ErrCalendaryDatesNoAvaliable
	}

	day := calendary[a.callCounter-1]

	switch day.DayType {
	case "playerOnSale":
		a.SetCurrentPlayerSale(nil, nil)
	case "playerSigning":
		a.SetCurrentPlayerSigning(nil)
	case "injury":
		a.SetCurrentInjuredPlayer(nil, nil)
	case "staffSigning":
		a.SetCurrentStaffSigning(nil)

	}
	log.Println("estamos en GetResume con day type", day.DayType)

	switch day.DayType {
	case "playerSigning":
		signings, err := a.ProcessPlayerSigning()
		if err != nil {
			log.Println("Error en la compra:", err)
			return []Calendary{}, err
		}
		log.Println("current signings is en GetResume de APP:", signings)
	case "playerOnSale":
		err := a.ProcessPlayerSale()
		if err != nil {
			log.Println("Error en la venta del jugador:", err)
			return []Calendary{}, err
		}
	case "injury":
		player, err := a.ProcessInjury()
		if err != nil {
			log.Println("Error APP get_resume, ProcessInjury", err)
			return []Calendary{}, err
		}
		log.Println("current injured player", player)

	case "staffSigning":
		staffSigning, err := a.ProcessStaffSigning()
		if err != nil {
			log.Println("Error APP get_resume, ProcessStaffSigning", err)
			return []Calendary{}, err
		}
		log.Println("current signings is en GetResume de APP:", staffSigning)

	case "staffSale":
		err := a.ProcessTeamStaffSale()
		if err != nil {
			log.Println("Error en la venta del Staff:", err)
			return []Calendary{}, err
		}
	case "match":
		if a.currentRivals == nil || len(*a.currentRivals) == 0 {
			rivals, err := a.rivalRepo.GetRival()
			if err != nil {
				log.Println("Error al obtener rivales:", err)
				return []Calendary{}, err
			}
			a.currentRivals = &rivals
		}
		if a.currentRivals == nil || len(*a.currentRivals) == 0 {
			log.Println("Error en el procesamiento del partido: currentRivals no está inicializado o está vacío")
			return []Calendary{}, ErrNoRivalsAvailable
		}

		match, err := a.ProcessMatch(day.CalendaryId)
		if err != nil {
			log.Println("Error en el procesamiento del partido:", err)
			return []Calendary{}, err
		}
		log.Println("partido actual", match)
	default:
		log.Println("Tipo de día desconocido:", day.DayType)
	}

	return []Calendary{day}, nil
}
