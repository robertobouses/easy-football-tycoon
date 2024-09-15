package app

import (
	"log"
)

func (a AppService) PostStrategy(req Strategy) error {

	log.Printf("Valores de player creado: %+v\n", req)
	err := a.strategyRepo.PostStrategy(req)
	if err != nil {
		log.Println("Error en PostStrategy:", err)
		return err
	}
	return nil
}
