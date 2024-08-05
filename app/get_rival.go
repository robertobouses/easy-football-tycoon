package app

import "log"

func (a AppService) GetRival() ([]Rival, error) {
	team, err := a.rivalRepo.GetRival()
	if err != nil {
		log.Println("Error al extraer GetRival", err)
		return []Rival{}, err
	}
	return team, nil
}
