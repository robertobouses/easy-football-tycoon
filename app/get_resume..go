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

	if day.DayType != "sale" {
		a.SetCurrentSalePlayer(nil)
	}

	if day.DayType != "purchase" {
		a.SetCurrentProspect(nil)
	}

	if day.DayType != "injury" {
		a.SetCurrentInjuredPlayer(nil)
	}

	log.Println("estasmos en GetResume con day type", day.DayType)

	switch day.DayType {
	case "purchase":
		prospect, err := a.ProcessPurchase()
		if err != nil {
			log.Println("Error en la compra:", err)
			return []Calendary{}, err
		}
		log.Println("current prospect is en GetResume de APP:", prospect)
	case "sale":
		err := a.ProcessSale()
		if err != nil {
			log.Println("Error en la venta:", err)
			return []Calendary{}, err
		}
	case "injury":
		player, err := a.ProcessInjury()
		if err != nil {
			log.Println("Error APP get_resume, ProcessInjury", err)
			return []Calendary{}, err
		}
		log.Println("current injured player", player)

	case "match":
		a.ProcessMatch()
	default:
		log.Println("Tipo de día desconocido:", day.DayType)
	}

	return []Calendary{day}, nil
}
