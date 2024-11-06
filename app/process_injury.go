package app

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func (a *AppService) ProcessInjury(playerId uuid.UUID) (Player, error) {
	var player Player
	var err error

	if playerId != uuid.Nil {
		playerPtr, err := a.teamRepo.GetPlayerByPlayerId(playerId)
		if err != nil {
			log.Println("Error en GetPlayerByPlayerId dentro de ProcessInjury:", err)
			return Player{}, err
		}
		player = *playerPtr
	} else {
		player, err = a.GetRandomPlayer()
		if err != nil {
			log.Println("Error al obtener un jugador aleatorio en ProcessInjury:", err)
			return Player{}, err
		}
	}
	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("Error al extraer GetAnalytics", err)
		return Player{}, err
	}

	injuryDays := a.calculateInjurySeverity(analytics.Physiotherapy)
	log.Println("injury Days son en ProcessInjury", injuryDays)
	a.SetCurrentInjuredPlayer(&player, &injuryDays)
	a.RunAutoPlayerDeclineByInjury(player.PlayerId, injuryDays)

	return player, nil

}

func (a *AppService) calculateInjurySeverity(physiotherapy int) int {
	rand.Seed(time.Now().UnixNano())
	injuryDays := WeightedRandom()

	switch {
	case physiotherapy >= 0 && physiotherapy < 20:
		return int(float64(injuryDays) * 2.4)
	case physiotherapy >= 20 && physiotherapy < 30:
		return injuryDays * 2
	case physiotherapy >= 30 && physiotherapy < 50:
		return int(float64(injuryDays) * 1.7)
	case physiotherapy >= 50 && physiotherapy < 70:
		return int(float64(injuryDays) * 1.3)
	case physiotherapy >= 70 && physiotherapy < 80:
		return injuryDays
	case physiotherapy >= 80 && physiotherapy < 90:
		return int(float64(injuryDays) * 0.9)
	case physiotherapy >= 90:
		return int(float64(injuryDays) * 0.8)
	default:
		log.Println("injuryDays son", injuryDays)
		return injuryDays

	}
}

func (a *AppService) GetCurrentInjuredPlayer() (*Player, *int, error) {
	return a.currentInjuredPlayer, a.injuryDays, nil
}
func (a *AppService) SetCurrentInjuredPlayer(player *Player, injuryDays *int) {
	a.currentInjuredPlayer = player
	a.injuryDays = injuryDays
}

func WeightedRandom() int {
	probablyInjuryDays := []int{
		1, 1, 1, 1,
		2, 2, 2, 2, 2, 2, 2,
		3, 3, 3, 3, 3, 3, 3, 3,
		4, 4, 4, 4, 4,
		5, 5, 5, 5,
		6, 6, 6,
		7, 7, 7,
		8, 8,
		9, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	}

	rand.Seed(time.Now().UnixNano())
	return probablyInjuryDays[rand.Intn(len(probablyInjuryDays))]
}
