package app

import (
	"log"
)

func (a AppService) PostRival(req Rival) error {

	log.Printf("Valores de player creado: %+v\n", req)
	err := a.rivalRepo.PostRival(req)
	if err != nil {
		log.Println("Error en PostRival:", err)
		return err
	}
	return nil
}
