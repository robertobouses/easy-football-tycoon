package app

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

func (a AppService) PostTeamStaff(req staff.Staff) error {

	log.Printf("Valores de player creado: %+v\n", req)
	err := a.teamStaffRepo.PostTeamStaff(req)
	if err != nil {
		log.Println("Error en PostStaff:", err)
		return err
	}
	return nil
}
