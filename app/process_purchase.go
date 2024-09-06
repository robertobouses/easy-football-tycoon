package app

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func (a *AppService) ProcessPurchase() (Prospect, error) {
	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("Error al extraer GetAnalytics", err)
		return Prospect{}, err
	}

	prospect, err := a.prospectRepo.GetProspectRandomByAnalytics(analytics.Scouting)
	log.Println("Prospecto asignado en ProcessPurchase 1:", prospect)
	if err != nil {
		log.Println("Error al extraer GetProspectRandomByAnalytics", err)
		return Prospect{}, err
	}

	if prospect.ProspectId == uuid.Nil {
		log.Println("No se encontró un prospecto válido")
		return Prospect{}, nil
	}

	a.SetCurrentProspect(&prospect)
	log.Println("Prospecto asignado en ProcessPurchase 2:", prospect)

	return prospect, nil
}

func (a *AppService) SetCurrentProspect(prospect *Prospect) {
	if prospect == nil || prospect.ProspectId == uuid.Nil {
		log.Println("Prospecto no válido, no se asignará. Prospecto:", prospect)
		a.currentProspect = nil
	} else {
		log.Println("Prospecto asignado en SetCurrentProspect:", *prospect)
		a.currentProspect = prospect
	}
}

func (a *AppService) GetCurrentProspect() (*Prospect, error) {
	log.Println("a.currentProspect en GetCurrentProspect 1:", a.currentProspect)
	if a.currentProspect == nil {
		log.Println("a.currentProspect es nil en GetCurrentProspect 2")
		return nil, nil
	}
	log.Println("Prospecto actual en GetCurrentProspect 3:", *a.currentProspect)
	return a.currentProspect, nil
}

func (a *AppService) AcceptPurchase(prospect *Prospect) error {

	initialFee := prospect.Fee

	paid, err := a.ProcessTransferFeePaid(initialFee)
	if err != nil {
		return ErrPaidNotFound
	}

	initalBalance, err := a.bankRepo.GetBalance()
	if err != nil {
		return ErrBalanceNotFound
	}

	newBalance := initalBalance - paid

	a.bankRepo.PostTransactions(paid, newBalance, prospect.ProspectName, "purchase")

	team := ConvertProspectToTeam(prospect)
	a.teamRepo.PostTeam(team)
	return nil
}

func (a *AppService) RejectPurchase(prospect *Prospect) {
	a.SetCurrentProspect(nil)
}

func ConvertProspectToTeam(prospect *Prospect) Team {
	return Team{
		PlayerId:   prospect.ProspectId,
		PlayerName: prospect.ProspectName,
		Position:   prospect.Position,
		Age:        prospect.Age,
		Fee:        prospect.Fee,
		Salary:     prospect.Salary,
		Technique:  prospect.Technique,
		Mental:     prospect.Mental,
		Physique:   prospect.Physique,
		InjuryDays: prospect.InjuryDays,
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
