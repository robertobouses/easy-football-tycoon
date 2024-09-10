package app

import (
	"log"
)

func (a AppService) PostTeamStaff(req Staff) error {

	log.Printf("Valores de player creado: %+v\n", req)
	err := a.teamStaffRepo.PostTeamStaff(req)
	if err != nil {
		log.Println("Error en PostStaff:", err)
		return err
	}
	return nil
}
