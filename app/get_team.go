package app

import "log"

func (a AppService) GetTeam() ([]Team, error) {
	team, err := a.lineupRepo.GetTeam()
	if err != nil {
		log.Println("Error al extraer crop", err)
		return []Team{}, err
	}
	return team, nil
}
