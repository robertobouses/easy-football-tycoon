package app

import "log"

func (a AppService) GetSignings() ([]Signings, error) {
	signings, err := a.signingsRepo.GetSignings()
	if err != nil {
		log.Println("Error al extraer GetSignings", err)
		return []Signings{}, err
	}
	return signings, nil
}
