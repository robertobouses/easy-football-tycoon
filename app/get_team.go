package app

import "log"

func (a AppService) GetTeam() ([]Player, error) {
	team, err := a.teamRepo.GetTeam()
	if err != nil {
		log.Println("Error al extraer GetTeam", err)
		return []Player{}, err
	}
	return team, nil
}
