package app

import (
	"log"
	"math/rand"
)

const (
	emergingPlayer    = 18
	veryYoungPlayer   = 21
	youngPlayer       = 24
	primePlayer       = 27
	experiencedPlayer = 31
	oldPlayer         = 33
	veryOldPlayer     = 35
)

func (a *AppService) RunAutoPlayerDevelopmentByAge(teamPlayers []Player) error {

	for _, player := range teamPlayers {
		var technique, mental, physique int
		switch {
		case player.Age < emergingPlayer:
			technique += ProbabilisticIncrement33() + ProbabilisticIncrement20() + ProbabilisticIncrement14()
			mental += ProbabilisticIncrement33() - ProbabilisticIncrement33() - ProbabilisticIncrement14() - ProbabilisticIncrement14()
			physique += ProbabilisticIncrement20() - ProbabilisticIncrement14()

		case player.Age < veryYoungPlayer:
			technique += rand.Intn(2) + ProbabilisticIncrement33() + ProbabilisticIncrement20() + ProbabilisticIncrement14() + ProbabilisticIncrement14()
			mental += ProbabilisticIncrement33() - ProbabilisticIncrement14() - ProbabilisticIncrement14()
			physique += rand.Intn(2) + ProbabilisticIncrement33() + ProbabilisticIncrement20() + ProbabilisticIncrement20() + ProbabilisticIncrement20() + ProbabilisticIncrement14() + ProbabilisticIncrement14()

		case player.Age < youngPlayer:
			technique += rand.Intn(4) + ProbabilisticIncrement33() + ProbabilisticIncrement33()
			mental += ProbabilisticIncrement66() + ProbabilisticIncrement20() - ProbabilisticIncrement33() - ProbabilisticIncrement20() - ProbabilisticIncrement14()
			physique += rand.Intn(4) + ProbabilisticIncrement66() + ProbabilisticIncrement66()

		case player.Age < primePlayer:
			technique += rand.Intn(4) + ProbabilisticIncrement33() + ProbabilisticIncrement33()
			mental += rand.Intn(2) - ProbabilisticIncrement20()
			physique -= ProbabilisticIncrement14()

		case player.Age < experiencedPlayer:
			technique += rand.Intn(2) - ProbabilisticIncrement20() - ProbabilisticIncrement20()
			mental += rand.Intn(3) + ProbabilisticIncrement33() + ProbabilisticIncrement20()
			physique -= ProbabilisticIncrement66() + ProbabilisticIncrement33()

		case player.Age < oldPlayer:
			technique += ProbabilisticIncrement66() + ProbabilisticIncrement14() + ProbabilisticIncrement14() - rand.Intn(2) - ProbabilisticIncrement20() - ProbabilisticIncrement14() - ProbabilisticIncrement14() - ProbabilisticIncrement14()
			mental += rand.Intn(4)
			physique -= rand.Intn(4) + ProbabilisticIncrement66() + ProbabilisticIncrement66() + ProbabilisticIncrement33()

		case player.Age < veryOldPlayer:
			technique -= rand.Intn(2) + ProbabilisticIncrement66() + ProbabilisticIncrement33()
			mental -= ProbabilisticIncrement33() + ProbabilisticIncrement14()
			physique -= rand.Intn(7) + ProbabilisticIncrement66() + ProbabilisticIncrement33() + ProbabilisticIncrement20() + ProbabilisticIncrement20()

		default:
			technique -= rand.Intn(3) + ProbabilisticIncrement66() + ProbabilisticIncrement66() + ProbabilisticIncrement14()
			mental -= rand.Intn(2) + ProbabilisticIncrement33() + ProbabilisticIncrement14()
			physique -= rand.Intn(8) + ProbabilisticIncrement66() + ProbabilisticIncrement66() + ProbabilisticIncrement33() + ProbabilisticIncrement20()
		}
		err := a.teamRepo.UpdatePlayerData(player.PlayerId, nil, nil, nil, nil, nil, nil, nil, &technique, &mental, &physique, nil, nil, nil, nil, nil)
		if err != nil {
			log.Printf("Error al actualizar los datos del jugador %v: %v", player.PlayerId, err)
		}

	}
	return nil
}
