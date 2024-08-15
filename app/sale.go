package app

import (
	"log"
	"math/rand"
	"time"
)

func (a AppService) Sale() (Team, error) {
	players, err := a.teamRepo.GetTeam()
	if err != nil {
		log.Println("Error al extraer GetTeam", err)
		return Team{}, err
	}

	if len(players) == 0 {
		log.Println("No se encontraron jugadores")
		return Team{}, nil
	}

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(players))
	player := players[randomIndex]

	return player, nil
}

//TODO ROBERTO CREO QUE ESTE MÉTODO DE TRAER Y LUEGO HACER ALEATORIO ES MÁS ÓPTIMO QUE EL QUE LO HACE TODO EN REPO GetProspectRandomByAnalytics
