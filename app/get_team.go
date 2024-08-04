package app

import "log"

func (a AppService) GetTeam() ([]Team, error) {
	team, err := a.teamRepo.GetTeam()
	if err != nil {
		log.Println("Error al extraer GetTeam", err)
		return []Team{}, err
	}
	return team, nil
}
