package app

import (
	"log"
	"math/rand"
)

const (
	masterTraining   = 93
	veryGoodTraining = 84
	goodTraining     = 69
	averageTraining  = 57
	basicTraining    = 48
	poorTraining     = 37
	badTraining      = 21
)

func (a *AppService) RunAutoPlayerDevelopmentByTraining(teamPlayers []Player) error {

	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("error al obtener analytics")
		return err
	}

	for _, player := range teamPlayers {
		technique := 0
		physique := 0
		mental := 0

		switch {
		case analytics.TotalTraining > masterTraining:
			technique += rand.Intn(7)
			mental += rand.Intn(2)
			physique += rand.Intn(5)

		case analytics.TotalTraining > veryGoodTraining:
			technique += rand.Intn(5)
			mental += ProbabilisticIncrement33()
			physique += rand.Intn(4)

		case analytics.TotalTraining > goodTraining:
			technique += rand.Intn(3)
			mental += ProbabilisticIncrement33()
			physique += rand.Intn(3) + ProbabilisticIncrement33()

		case analytics.TotalTraining > averageTraining:
			technique += rand.Intn(2)
			mental += ProbabilisticIncrement14()
			physique += rand.Intn(3) + ProbabilisticIncrement14()

		case analytics.TotalTraining > basicTraining:
			technique = rand.Intn(2)
			physique += ProbabilisticIncrement66() + ProbabilisticIncrement14()

		case analytics.TotalTraining > poorTraining:
			technique = ProbabilisticIncrement33()
			physique += rand.Intn(2) + ProbabilisticIncrement14()

		case analytics.TotalTraining > badTraining:
			technique = ProbabilisticIncrement14()
			physique += ProbabilisticIncrement33()

		default:
			physique += ProbabilisticIncrement20()
		}

		err = a.teamRepo.UpdatePlayerData(player.PlayerId, nil, nil, nil, nil, nil, nil, nil, &technique, &mental, &physique, nil, nil, nil, nil, nil)
		if err != nil {
			log.Printf("Error al actualizar los datos del jugador %v: %v", player.PlayerId, err)
		}
	}
	return nil
}
