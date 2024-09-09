package app

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

func (a *AppService) ProcessPlayerSale() error {
	player, err := a.GetRandomPlayer()
	if err != nil {
		return err
	}
	if player == (Team{}) {
		return errors.New("no player available for sale")
	}

	transferFeeReceived, err := a.ProcessTransferFeeReceived(player.Fee)
	if err != nil {
		return err
	}

	a.SetCurrentSalePlayer(&player, &transferFeeReceived)

	return nil
}

func (a *AppService) GetCurrentSalePlayer() (*Team, *int, error) {
	return a.currentPlayerOnSale, a.transferFeeReceived, nil
}

func (a *AppService) SetCurrentSalePlayer(player *Team, transferFeeReceived *int) {
	a.currentPlayerOnSale = player
	a.transferFeeReceived = transferFeeReceived
}

func (a *AppService) AcceptPlayerSale(player Team) error {
	if a.transferFeeReceived == nil {
		return ErrTransferNotFound
	}

	balance, err := a.bankRepo.GetBalance()
	if err != nil {
		return ErrBalanceNotFound
	}

	amount := *a.transferFeeReceived
	newBalance := balance + amount
	log.Println("amount en APP", amount)
	log.Println("balance inicial en APP", balance)
	log.Println("new balance en APP", newBalance)

	err = a.bankRepo.PostTransactions(amount, newBalance, player.PlayerName, "sale")
	if err != nil {
		return err
	}

	err = a.teamRepo.DeletePlayerFromTeam(player)
	if err != nil {
		return err
	}

	return nil
}

func (a *AppService) RejectPlayerSale(player Team) {
	a.SetCurrentSalePlayer(nil, nil)
}

// TODO ROBERTO CREO QUE ESTE MÉTODO DE TRAER Y LUEGO HACER ALEATORIO ES MÁS ÓPTIMO QUE EL QUE LO HACE TODO EN REPO GetSigningsRandomByAnalytics
func (a *AppService) GetRandomPlayer() (Team, error) {
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

func (a *AppService) ProcessTransferFeeReceived(initialFee int) (int, error) {
	rand.Seed(time.Now().UnixNano())
	randomFactor := 0.2 + rand.Float64()*2.3
	randomFee := float64(initialFee) * randomFactor

	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		return 0, err
	}

	var finalFee float64
	switch {
	case analytics.Finances <= 10:
		finalFee = randomFee * 0.5
	case analytics.Finances <= 30:
		finalFee = randomFee * 0.8
	case analytics.Finances <= 50:
		finalFee = randomFee
	case analytics.Finances <= 70:
		finalFee = randomFee * 1.2
	case analytics.Finances <= 80:
		finalFee = randomFee * 1.3
	case analytics.Finances <= 90:
		finalFee = randomFee * 1.4
	default:
		finalFee = randomFee * 1.5
	}

	return int(finalFee), nil
}
