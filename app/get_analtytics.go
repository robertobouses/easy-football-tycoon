package app

import "log"

func (a AppService) GetAnalytics() (Analytics, error) {
	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("Error al extraer GetAnalytics", err)
		return Analytics{}, err
	}
	return analytics, nil
}
