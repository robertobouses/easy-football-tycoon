package app

import (
	"log"
	"math/rand"
	"time"
)

func (a AppService) PostCalendary() error {

	calendary := generateCalendary()

	log.Printf("Valores de player creado: %+v\n", calendary)
	err := a.calendaryRepo.PostCalendary(calendary)
	if err != nil {
		log.Println("Error en PostCalendary:", err)
		return err
	}
	return nil
}

func shuffle(slice []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func generateCalendary() []string {

	primerSegmento := []string{"playerSigning", "playerSigning", "playerSigning", "playerSigning", "playerSigning", "staffSigning", "staffSigning", "staffSigning",
		"playerOnSale", "playerOnSale", "playerOnSale", "playerOnSale", "playerOnSale", "staffSale",
		"injury", "injury"}
	segundoSegmento := []string{"match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match",
		"injury", "injury", "injury"}
	tercerSegmento := []string{"match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match", "match",
		"injury", "injury", "injury"}
	cuartoSegmento := []string{"playerSigning", "playerSigning", "playerSigning", "playerSigning", "playerSigning", "playerSigning", "staffSigning", "staffSigning",
		"staffSigning", "playerOnSale", "playerOnSale", "playerOnSale", "staffSale", "staffSale",
		"injury", "injury"}

	shuffle(primerSegmento)
	shuffle(segundoSegmento)
	shuffle(tercerSegmento)
	shuffle(cuartoSegmento)

	calendary := append(primerSegmento, segundoSegmento...)
	calendary = append(calendary, tercerSegmento...)
	calendary = append(calendary, cuartoSegmento...)

	return calendary
}
