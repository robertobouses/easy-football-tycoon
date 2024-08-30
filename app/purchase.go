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
	return prospect, nil
}
