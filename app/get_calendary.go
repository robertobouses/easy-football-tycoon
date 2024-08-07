package app

import "log"

func (a AppService) GetCalendary() ([]Calendary, error) {
	calendary, err := a.calendaryRepo.GetCalendary()
	if err != nil {
		log.Println("Error al extraer GetCalendary", err)
		return []Calendary{}, err
	}
	return calendary, nil
}
