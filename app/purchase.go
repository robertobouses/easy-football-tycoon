package app

import "log"

func (a AppService) Purchase() {
	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("Error al extraer GetAnalytics", err)
		return Analytics{}, err
	}

	prospect, err := a.prospectRepo.GetProspectRandomByAnalytics(analytics.scouting)
	if err != nil {
		log.Println("Error al extraer GetProspect", err)
		return Prospect{}, err
	}
	return prospect, nil
}
