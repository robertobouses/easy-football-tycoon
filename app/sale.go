package app

import (
	"log"
	"math/rand"
	"time"
)

func (a AppService) GetRandomPlayerForSale() (Team, error) {
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

func (a AppService) Sale() error {
	player, err := a.GetRandomPlayerForSale()
	if err != nil {
		return err
	}
	a.currentSalePlayer = &player

	return nil
}
func (a AppService) AcceptSale(player Team) error {
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

func (a AppService) RejectSale(player Team) {

}

// TODO ROBERTO CREO QUE ESTE MÉTODO DE TRAER Y LUEGO HACER ALEATORIO ES MÁS ÓPTIMO QUE EL QUE LO HACE TODO EN REPO GetProspectRandomByAnalytics
func (a *AppService) GetCurrentSalePlayer() *Team {
	return a.currentSalePlayer
}

func (a *AppService) SetCurrentSalePlayer(player *Team) {
	a.currentSalePlayer = player
}
