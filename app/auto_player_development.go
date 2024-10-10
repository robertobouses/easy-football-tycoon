package app

import (
	"log"
	"math/rand"
)

const (
	masterTraining   = 92
	veryGoodTraining = 83
	goodTraining     = 69
	averageTraining  = 57
	basicTraining    = 48
	poorTraining     = 37
	badTraining      = 21
)

func (a *AppService) RunAutoPlayerDevelopment() error {
	//al acabar año cambio age y cambio valor y salary decidir si se negocia renovacion
	//TODO GESTION ERRORES EN FORMULAS AUTOCALCULO DE JUGADORES Y STAFF
	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("error al obtener analytics")
		return err
	}

	teamPlayers, err := a.teamRepo.GetTeam()
	if err != nil {
		log.Println("error al obtener team players")
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
			physique += rand.Intn(6)

		case analytics.TotalTraining > veryGoodTraining:
			technique += rand.Intn(5)
			mental += probabilisticIncrement33()
			physique += rand.Intn(4)

		case analytics.TotalTraining > goodTraining:
			technique += rand.Intn(3)
			mental += probabilisticIncrement33()
			physique += rand.Intn(3) + probabilisticIncrement33()

		case analytics.TotalTraining > averageTraining:
			technique += rand.Intn(2)
			mental += probabilisticIncrement14()
			physique += rand.Intn(3) + probabilisticIncrement14()

		case analytics.TotalTraining > basicTraining:
			technique = rand.Intn(2)
			mental = 0
			physique += probabilisticIncrement66() + probabilisticIncrement14()

		case analytics.TotalTraining > poorTraining:
			technique = probabilisticIncrement33()
			mental = 0
			physique += rand.Intn(2) + probabilisticIncrement14()

		case analytics.TotalTraining > badTraining:
			technique = probabilisticIncrement14()
			mental = 0
			physique += probabilisticIncrement33()

		default:
			technique = 0
			mental = 0
			physique += probabilisticIncrement20()
		}

		err = a.teamRepo.UpdatePlayerData(player.PlayerId, nil, nil, nil, nil, nil, nil, nil, &technique, &mental, &physique, nil, nil, nil, nil, nil)
		if err != nil {
			log.Printf("Error al actualizar los datos del jugador %v: %v", player.PlayerId, err)
		}
	}
	return nil
}

func probabilisticIncrement14() int {
	if rand.Intn(7) == 4 {
		return 1
	}
	return 0
}

func probabilisticIncrement20() int {
	if rand.Intn(5) == 4 {
		return 1
	}
	return 0
}

func probabilisticIncrement33() int {
	if rand.Intn(3) == 2 {
		return 1
	}
	return 0
}

func probabilisticIncrement66() int {
	if rand.Intn(3) < 2 {
		return 1
	}
	return 0
}
