package app

import (
	"log"
)

func (a AppService) PostAnalytics(req Analytics) error {

	log.Printf("Valores de player creado: %+v\n", req)
	err := a.analyticsRepo.PostAnalytics(req)
	if err != nil {
		log.Println("Error en PostAnalytics:", err)
		return err
	}
	return nil
}
