package app

import "log"

func (a AppService) GetResume() ([]Calendary, error) {
	calendary, err := a.calendaryRepo.GetCalendary()
	if err != nil {
		log.Println("Error al extraer GetCalendary:", err)
		return []Calendary{}, err
	}

	for _, day := range calendary {
		switch day.DayType {
		case "purchase":
			a.Purchase()
		case "sale":
			a.Sale()
		case "injury":
			a.Injury()
		case "match":
			a.Match()
		default:
			log.Println("Tipo de día desconocido:", day.Daytype)
		}
	}

	return calendary, nil
}
