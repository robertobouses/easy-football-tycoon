package app

import "log"

func (a AppService) GetProspect() ([]Prospect, error) {
	prospect, err := a.prospectRepo.GetProspect()
	if err != nil {
		log.Println("Error al extraer GetProspect", err)
		return []Prospect{}, err
	}
	return prospect, nil
}
