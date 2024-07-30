package app

import "log"

func (a AppService) GetLineup() ([]Lineup, error) {
	lineup, err := a.lineupRepo.GetLineup()
	if err != nil {
		log.Println("Error al extraer crop", err)
		return []Lineup{}, err
	}
	return lineup, nil
}
