package app

import (
	"log"
)

func (a AppService) PostProspect(req Prospect) error {

	log.Printf("Valores de player creado: %+v\n", req)
	err := a.prospectRepo.PostProspect(req)
	if err != nil {
		log.Println("Error en PostProspect:", err)
		return err
	}
	return nil
}
