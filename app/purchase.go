package app

import "log"

func (a AppService) Purchase() (Prospect, error) {
	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("Error al extraer GetAnalytics", err)
		return Prospect{}, err
	}

	prospect, err := a.prospectRepo.GetProspectRandomByAnalytics(analytics.Scouting)
	if err != nil {
		log.Println("Error al extraer GetProspectRandomByAnalytics", err)
		return Prospect{}, err
	}
	a.SetCurrentProspect(&prospect)

	return prospect, nil
}

func (a *AppService) GetCurrentProspect() (*Prospect, error) {
	return a.currentProspect, nil
}

func (a *AppService) SetCurrentProspect(prospect *Prospect) {
	a.currentProspect = prospect
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
