package app

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

func (a AppService) GetTeamStaff() ([]staff.Staff, error) {
	staffs, err := a.teamStaffRepo.GetTeamStaff()
	if err != nil {
		log.Println("Error al extraer GetTeam", err)
		return []staff.Staff{}, err
	}
	return staffs, nil
}
