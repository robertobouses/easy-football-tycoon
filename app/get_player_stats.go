package app

import "log"

func (a *AppService) GetPlayerStats() ([]PlayerStats, error) {
	stats, err := a.statsRepo.GetPlayerStats()
	if err != nil {
		log.Println("Error al extraer GetPlayerStats", err)
		return []PlayerStats{}, err
	}
	return stats, nil
}
