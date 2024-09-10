package app

import "log"

func (a AppService) GetTeamStaff() ([]Staff, error) {
	staff, err := a.teamStaffRepo.GetTeamStaff()
	if err != nil {
		log.Println("Error al extraer GetTeam", err)
		return []Staff{}, err
	}
	return staff, nil
}
