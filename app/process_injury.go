package app

import (
	"log"
	"math/rand"
	"time"
)

func (a *AppService) ProcessInjury() (Team, error) {
	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("Error al extraer GetAnalytics", err)
		return Team{}, err
	}
	player, err := a.GetRandomPlayer()
	if err != nil {
		return Team{}, err
	}

	injuryDays := a.calculateInjurySeverity(analytics.Physiotherapy)
	log.Println("injury Days son en ProcessInjury", injuryDays)
	a.SetCurrentInjuredPlayer(&player, &injuryDays)

	lined := false
	a.teamRepo.UpdatePlayerData(player.PlayerId, nil, nil, nil, nil, nil, nil, nil, nil, &injuryDays, &lined)

	return player, nil
}

func (a *AppService) calculateInjurySeverity(physiotherapy int) int {
	rand.Seed(time.Now().UnixNano())
	injuryDays := rand.Intn(29) + 2

	switch {
	case physiotherapy >= 0 && physiotherapy < 20:
		return int(float64(injuryDays) * 2.4)
	case physiotherapy >= 20 && physiotherapy < 30:
		return injuryDays * 2
	case physiotherapy >= 30 && physiotherapy < 50:
		return int(float64(injuryDays) * 1.8)
	case physiotherapy >= 50 && physiotherapy < 70:
		return int(float64(injuryDays) * 1.5)
	case physiotherapy >= 70 && physiotherapy < 80:
		return int(float64(injuryDays) * 1.2)
	case physiotherapy >= 80 && physiotherapy < 90:
		return injuryDays
	case physiotherapy >= 90:
		return int(float64(injuryDays) * 0.8)
	default:
		log.Println("injuryDays son", injuryDays)
		return injuryDays

	}
}

func (a *AppService) GetCurrentInjuredPlayer() (*Team, *int, error) {
	return a.currentInjuredPlayer, a.injuryDays, nil
}
func (a *AppService) SetCurrentInjuredPlayer(player *Team, injuryDays *int) {
	a.currentInjuredPlayer = player
	a.injuryDays = injuryDays
}
