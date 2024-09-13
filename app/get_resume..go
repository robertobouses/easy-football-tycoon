package app

import (
	"log"
)

func (a *AppService) GetResume() ([]Calendary, error) {

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
