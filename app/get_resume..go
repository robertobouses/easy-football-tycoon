package app

import "log"

func (a AppService) GetResume() error {
	calendary, err := a.calendaryRepo.GetCalendary()
	if err != nil {
		log.Println("Error al extraer GetCalendary", err)
		return []Calendary{}, err
	}
	return calendary, nil

	if calendary.Daytype == "purchase" {
		a.Purchase()
	}
	if calendary.Daytype == "sale" {
		a.Sale()
	}
	if calendary.Daytype == "injury" {
		a.Injury()
	}
	if calendary.Daytype == "match" {
		a.Match()
	}

}
