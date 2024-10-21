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

	var goalkeeperCount, defenderCount, midfieldersCount, forwardCount int

	//TODO COMPROBACIONES DE ALINEACION EN UNA FUNCIÓN FUERA DE GET RESUME

	for _, player := range lineup {
		switch player.Position {
		case "goalkeeper":
			goalkeeperCount++
		case "defender":
			defenderCount++
		case "midfielder":
			midfieldersCount++
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

		if goalkeeperCount != 1 || defenderCount != 4 || midfieldersCount != 4 || forwardCount != 2 {
			log.Println("La alineación no cumple con la formación 4-4-2.")
			return []Calendary{}, ErrInvalidFormation
		}
	case "4-3-3":

		if goalkeeperCount != 1 || defenderCount != 4 || midfieldersCount != 3 || forwardCount != 3 {
			log.Println("La alineación no cumple con la formación 4-3-3.")
			return []Calendary{}, ErrInvalidFormation
		}

	case "4-5-1":

		if goalkeeperCount != 1 || defenderCount != 4 || midfieldersCount != 5 || forwardCount != 1 {
			log.Println("La alineación no cumple con la formación 4-5-1.")
			return []Calendary{}, ErrInvalidFormation
		}
	case "5-3-2":

		if goalkeeperCount != 1 || defenderCount != 5 || midfieldersCount != 3 || forwardCount != 2 {
			log.Println("La alineación no cumple con la formación 5-3-2.")
			return []Calendary{}, ErrInvalidFormation
		}
	case "3-4-3":

		if goalkeeperCount != 1 || defenderCount != 3 || midfieldersCount != 4 || forwardCount != 3 {
			log.Println("La alineación no cumple con la formación 3-4-3.")
			return []Calendary{}, ErrInvalidFormation
		}
	case "3-5-2":

		if goalkeeperCount != 1 || defenderCount != 3 || midfieldersCount != 5 || forwardCount != 2 {
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

	if a.callCounter >= len(calendary) {
		log.Println("No hay más fechas disponibles en el calendario")
		return []Calendary{}, ErrCalendaryDatesNoAvailable
	}

	lenCalendary := len(calendary)
	i := lenCalendary / 3
	j := (2 * lenCalendary) / 3
	k := lenCalendary - 1

	team, err := a.teamRepo.GetTeam()
	if err != nil {
		log.Println("Error obtener Team en GetResume:", err)
		return nil, err
	}

	staff, err := a.staffRepo.GetStaff()
	if err != nil {
		log.Println("Error obtener Staff en GetResume:", err)
		return nil, err
	}

	thirdOfPlayersSalary := calculateTotalTeamSalary(team) / 3
	thirdOfStaffSalary := calculateTotalStaffSalary(staff) / 3

	if a.callCounter == i || a.callCounter == j || a.callCounter == k {
		err := a.RunAutoPlayerDevelopmentByTraining(team)
		if err != nil {
			log.Println("Error al automatizarel desarrollo de los jugadores por entrenamiento")
		}

		err = a.RunAutoPlayerDevelopmentByAge(team)
		if err != nil {
			log.Println("Error al automatizar el desarrollo de los jugadores por edad")
		}
		initialBalance, err := a.bankRepo.GetBalance()
		if err != nil {
			log.Println("Error al obtener el balance inicial:", err)
			return nil, err
		}
		newBalance := initialBalance - thirdOfPlayersSalary

		err = a.bankRepo.PostTransactions(thirdOfPlayersSalary, newBalance, "Player Salaries Payment", "Salaries")
		if err != nil {
			log.Println("Error al procesar el pago de salarios de los jugadores:", err)
			return nil, err
		}

		initialBalance, err = a.bankRepo.GetBalance()
		if err != nil {
			log.Println("Error al obtener el balance inicial:", err)
			return nil, err
		}
		newBalance = initialBalance - thirdOfStaffSalary
		err = a.bankRepo.PostTransactions(thirdOfStaffSalary, newBalance, "Staff Salaries Payment", "Salaries")
		if err != nil {
			log.Println("Error al procesar el pago de salarios del personal:", err)
			return nil, err
		}

		log.Println("Salarios pagados exitosamente en la fecha del calendario:", a.callCounter)
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

		// match, err := a.ProcessMatchSimulation(day.CalendaryId)
		// if err != nil {
		// 	log.Println("Error en el procesamiento del partido:", err)
		// 	return []Calendary{}, err
		// }

		// 	log.Println("partido actual", match)
		// default:
		// 	log.Println("Tipo de día desconocido:", day.DayType)
		// }

	}
	log.Println("Día de partido encontrado. Esperando decisión de jugar o simular.")
	return []Calendary{day}, nil
}
