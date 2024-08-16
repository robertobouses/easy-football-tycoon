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
			err := a.Sale()
			if err != nil {
				log.Println("Error en la venta:", err)
				return []Calendary{}, err
			}
		case "injury":
			a.Injury()
		case "match":
			a.Match()
		default:
			log.Println("Tipo de día desconocido:", day.DayType)
		}
	}

	return calendary, nil
}
