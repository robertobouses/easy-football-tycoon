package app

import (
	"log"

	"github.com/google/uuid"
)

func (a AppService) ProcessPurchase() (Prospect, error) {
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
