package app

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/easy-football-tycoon/app/signings"
)

func (a *AppService) ProcessPlayerSigning() (signings.Signings, error) {
	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("Error al extraer GetAnalytics", err)
		return signings.Signings{}, err
	}

	signing, err := a.signingsRepo.GetSigningsRandomByAnalytics(analytics.Scouting)
	log.Println("Signings asignado en ProcessPlayerSigning 1:", signing)
	if err != nil {
		log.Println("Error al extraer GetSigningsRandomByAnalytics", err)
		return signings.Signings{}, err
	}

	//TODO ESTABLECER UN FACTOR PARA QUE SEA RAMDON EL PRECIO DEL FICHAJES, RELACIONADO CON FINANCES
	//transferFeeIssued utilizar función ProcessTransferFeeReceived

	if signing.SigningsId == uuid.Nil {
		log.Println("No se encontró un signings válido")
		return signings.Signings{}, nil
	}

	a.SetCurrentPlayerSigning(&signing)
	log.Println("Signings asignado en ProcessPlayerSigning 2:", signing)

	return signing, nil
}

func (a *AppService) SetCurrentPlayerSigning(signings *signings.Signings) {
	if signings == nil || signings.SigningsId == uuid.Nil {
		log.Println("Signings no válido, no se asignará. Signings:", signings)
		a.currentPlayerSigning = nil
	} else {
		log.Println("Signings asignado en SetCurrentPlayerSigning:", *signings)
		a.currentPlayerSigning = signings
	}
}

func (a *AppService) GetCurrentPlayerSigning() (*signings.Signings, error) {
	log.Println("a.currentPlayerSigning en GetCurrentPlayerSigning 1:", a.currentPlayerSigning)
	if a.currentPlayerSigning == nil {
		log.Println("a.currentPlayerSigning es nil en GetCurrentPlayerSigning 2")
		return nil, nil
	}
	log.Println("Signings actual en GetCurrentPlayerSigning 3:", *a.currentPlayerSigning)
	return a.currentPlayerSigning, nil
}

func (a *AppService) AcceptPlayerSigning(signings *signings.Signings) error {

	initialFee := signings.Fee

	paid, err := a.ProcessTransferFeePaid(initialFee)
	if err != nil {
		return ErrPaidNotFound
	}

	initialBalance, err := a.bankRepo.GetBalance()
	if err != nil {
		return ErrBalanceNotFound
	}

	newBalance := initialBalance - paid

	a.bankRepo.PostTransactions(paid, newBalance, signings.LastName, "Player Signing")

	team := ConvertSigningsToTeam(signings)
	a.teamRepo.PostTeam(team)

	err = a.signingsRepo.DeleteSigning(*signings)
	if err != nil {
		return err
	}

	return nil
}

func (a *AppService) RejectPlayerSigning(signings *signings.Signings) {
	a.SetCurrentPlayerSigning(nil)
}

func ConvertSigningsToTeam(signings *signings.Signings) Player {
	randomFactor := 0.8 + rand.Float64()*1.5
	happiness := int(70 * randomFactor)
	happiness = int(math.Min(float64(happiness), 100))

	randomFactor2 := 1 + rand.Float64()*17
	familiarity := int(randomFactor2)

	return Player{
		PlayerId:    signings.SigningsId,
		FirstName:   signings.FirstName,
		LastName:    signings.LastName,
		Nationality: signings.Nationality,
		Position:    signings.Position,
		Age:         signings.Age,
		Fee:         signings.Fee,
		Salary:      signings.Salary,
		Technique:   signings.Technique,
		Mental:      signings.Mental,
		Physique:    signings.Physique,
		InjuryDays:  signings.InjuryDays,
		Lined:       false,
		Familiarity: familiarity,
		Fitness:     signings.Fitness,
		Happiness:   happiness,
	}
}

func (a *AppService) ProcessTransferFeePaid(initialFee int) (int, error) {
	rand.Seed(time.Now().UnixNano())
	randomFactor := 0.8 + rand.Float64()*1.9
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
