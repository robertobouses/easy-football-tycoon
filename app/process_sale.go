package app

import (
	"log"
	"math/rand"
	"time"
)

func (a *AppService) ProcessSale() error {
	player, err := a.GetRandomPlayer()
	if err != nil {
		return err
	}
	a.SetCurrentSalePlayer(&player)

	return nil
}

func (a *AppService) GetCurrentSalePlayer() (*Team, error) {
	return a.currentSalePlayer, nil
}

func (a *AppService) SetCurrentSalePlayer(player *Team) {
	a.currentSalePlayer = player
}

func (a *AppService) AcceptSale(player Team) error {
	err := a.teamRepo.DeletePlayerFromTeam(player)
	if err != nil {
		return err
	}
	//TODO PARTE VENTA FINANZAS TRAER NIVEL ANALYTICS FINANCE Y EJECUTAR CONTRA SALDO DB CUENTA BANCARIA

	// 	err = a.moneyRepo.AddMoney(player.Value)
	// 	if err != nil {
	// 		return err
	// 	}

	return nil
}

func (a *AppService) RejectSale(player Team) {
	a.SetCurrentSalePlayer(nil)
}

// TODO ROBERTO CREO QUE ESTE MÉTODO DE TRAER Y LUEGO HACER ALEATORIO ES MÁS ÓPTIMO QUE EL QUE LO HACE TODO EN REPO GetProspectRandomByAnalytics
func (a AppService) GetRandomPlayer() (Team, error) {
	players, err := a.teamRepo.GetTeam()
	if err != nil {
		log.Println("Error al extraer GetTeam:", err)
		return Team{}, err
	}

	if len(players) == 0 {
		log.Println("No se encontraron jugadores")
		return Team{}, nil
	}

	log.Printf("Jugadores obtenidos: %+v", players)

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(players))
	player := players[randomIndex]

	log.Printf("Jugador seleccionado aleatoriamente para la venta: %+v", player)

	return player, nil
}
