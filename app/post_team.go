package app

import (
	"log"
)

func (a AppService) PostTeam(req Player) error {

	log.Printf("Valores de player creado: %+v\n", req)
	err := a.teamRepo.PostTeam(req)
	if err != nil {
		log.Println("Error en PostTeam:", err)
		return err
	}
	return nil
}
