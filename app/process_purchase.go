package app

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func (a *AppService) ProcessPurchase() (Signings, error) {
	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("Error al extraer GetAnalytics", err)
		return Signings{}, err
	}

	signings, err := a.signingsRepo.GetSigningsRandomByAnalytics(analytics.Scouting)
	log.Println("Signingso asignado en ProcessPurchase 1:", signings)
	if err != nil {
		log.Println("Error al extraer GetSigningsRandomByAnalytics", err)
		return Signings{}, err
	}

	if signings.SigningsId == uuid.Nil {
		log.Println("No se encontró un signingso válido")
		return Signings{}, nil
	}

	a.SetCurrentSignings(&signings)
	log.Println("Signingso asignado en ProcessPurchase 2:", signings)

	return signings, nil
}

func (a *AppService) SetCurrentSignings(signings *Signings) {
	if signings == nil || signings.SigningsId == uuid.Nil {
		log.Println("Signingso no válido, no se asignará. Signingso:", signings)
		a.currentSignings = nil
	} else {
		log.Println("Signingso asignado en SetCurrentSignings:", *signings)
		a.currentSignings = signings
	}
}

func (a *AppService) GetCurrentSignings() (*Signings, error) {
	log.Println("a.currentSignings en GetCurrentSignings 1:", a.currentSignings)
	if a.currentSignings == nil {
		log.Println("a.currentSignings es nil en GetCurrentSignings 2")
		return nil, nil
	}
	log.Println("Signingso actual en GetCurrentSignings 3:", *a.currentSignings)
	return a.currentSignings, nil
}

func (a *AppService) AcceptPurchase(signings *Signings) error {

	initialFee := signings.Fee

	paid, err := a.ProcessTransferFeePaid(initialFee)
	if err != nil {
		return ErrPaidNotFound
	}

	initalBalance, err := a.bankRepo.GetBalance()
	if err != nil {
		return ErrBalanceNotFound
	}

	newBalance := initalBalance - paid

	a.bankRepo.PostTransactions(paid, newBalance, signings.SigningsName, "purchase")

	team := ConvertSigningsToTeam(signings)
	a.teamRepo.PostTeam(team)
	return nil
}

func (a *AppService) RejectPurchase(signings *Signings) {
	a.SetCurrentSignings(nil)
}

func ConvertSigningsToTeam(signings *Signings) Team {
	return Team{
		PlayerId:   signings.SigningsId,
		PlayerName: signings.SigningsName,
		Position:   signings.Position,
		Age:        signings.Age,
		Fee:        signings.Fee,
		Salary:     signings.Salary,
		Technique:  signings.Technique,
		Mental:     signings.Mental,
		Physique:   signings.Physique,
		InjuryDays: signings.InjuryDays,
		Lined:      false,
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
