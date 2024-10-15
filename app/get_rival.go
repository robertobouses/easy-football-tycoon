package app

import "log"

func (a AppService) GetRival() ([]Rival, error) {
	teams, err := a.rivalRepo.GetRival()
	if err != nil {
		log.Println("Error al extraer GetRival", err)
		return []Rival{}, err
	}
	return teams, nil
}
