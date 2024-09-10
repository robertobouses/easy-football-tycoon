package app

import (
	"log"
)

func (a AppService) PostSignings(req Signings) error {

	log.Printf("Valores de player creado: %+v\n", req)
	err := a.signingsRepo.PostSignings(req)
	if err != nil {
		log.Println("Error en PostSignings:", err)
		return err
	}
	return nil
}
