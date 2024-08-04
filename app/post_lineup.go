package app

import (
	"log"
)

func (a AppService) PostLineup(req Lineup) error {

	id := req.PlayerId

	log.Printf("Valores de player creado: %+v\n", req)
	err := a.lineupRepo.PostLineup(id)
	if err != nil {
		log.Println("Error en PostLineup:", err)
		return err
	}
	err = a.teamRepo.UpdatePlayerLinedStatus(id)
	if err != nil {
		log.Println("Error en UpdatePlayerLinedStatus:", err)
		return err
	}
	return nil
}
