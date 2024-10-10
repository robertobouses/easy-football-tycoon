package app

import (
	"log"
	"math/rand"

	"github.com/google/uuid"
)

const (
	briefTime     = 4
	shortTime     = 7
	mediumTime    = 12
	longTime      = 19
	extendedTime  = 27
	prolongedTime = 43
)

func (a *AppService) RunAutoPlayerDeclineByInjury(playerId uuid.UUID, injuryDays int) error {
	var technique, mental, physique, familiarity, fitness, happiness, fee int
	switch {
	case injuryDays < briefTime:
		technique -= rand.Intn(2)
		physique -= rand.Intn(3) + 1
		familiarity -= rand.Intn(10) + 1
		fitness -= rand.Intn(20) + 25
		happiness -= rand.Intn(14)

	case injuryDays < shortTime:
		technique -= rand.Intn(3)
		mental -= rand.Intn(2) - ProbabilisticSelfImprovement3()
		physique -= rand.Intn(3) + 1
		familiarity -= rand.Intn(12) + 2
		fitness -= rand.Intn(20) + 30
		happiness -= rand.Intn(20)
		fee -= 20_000

	case injuryDays < mediumTime:
		technique -= rand.Intn(3) + 1
		mental -= rand.Intn(2) - ProbabilisticSelfImprovement3()
		physique -= rand.Intn(5) + 3
		familiarity -= rand.Intn(12) + 10
		fitness -= rand.Intn(20) + 35
		happiness -= rand.Intn(30)

		decreasePercentage := rand.Intn(7) + 3
		fee -= (fee * decreasePercentage) / 100

	case injuryDays < longTime:
		technique -= rand.Intn(5) + 2
		mental -= rand.Intn(4) - ProbabilisticSelfImprovement3()
		physique -= rand.Intn(7) + 4
		familiarity -= rand.Intn(15) + 15
		fitness -= rand.Intn(26) + 40
		happiness -= rand.Intn(30) + 10

		decreasePercentage := rand.Intn(10) + 7
		fee -= (fee * decreasePercentage) / 100

	case injuryDays < extendedTime:
		technique -= rand.Intn(7) + 2
		mental -= rand.Intn(5) - ProbabilisticSelfImprovement8()
		physique -= rand.Intn(8) + 7
		familiarity -= rand.Intn(20) + 15
		fitness -= rand.Intn(26) + 60
		happiness -= rand.Intn(35) + 10

		decreasePercentage := rand.Intn(12) + 10
		fee -= (fee * decreasePercentage) / 100

	case injuryDays < prolongedTime:
		technique -= rand.Intn(8) + 3
		mental -= rand.Intn(5) - ProbabilisticSelfImprovement8()
		physique -= rand.Intn(8) + 12
		familiarity -= rand.Intn(20) + 18
		fitness -= rand.Intn(31) + 70
		happiness -= rand.Intn(40) + 10

		decreasePercentage := rand.Intn(18) + 15
		fee -= (fee * decreasePercentage) / 100

	default:
		technique -= rand.Intn(12) + 4
		mental -= rand.Intn(7) - ProbabilisticSelfImprovement8()
		physique -= rand.Intn(11) + 18
		familiarity -= rand.Intn(30) + 20
		fitness -= rand.Intn(26) + 75
		happiness -= rand.Intn(40) + 30

		fee -= (fee * 40) / 100

	}
	err := a.teamRepo.UpdatePlayerData(playerId, nil, nil, nil, nil, nil, &fee, nil, &technique, &mental, &physique, nil, nil, &familiarity, &fitness, &happiness)
	if err != nil {
		log.Println("error al calcular el declive del jugador lesionado")
	}
	return nil

}

func ProbabilisticSelfImprovement8() int {
	if rand.Intn(3) < 2 {
		return 1
	}
	return 8

}

func ProbabilisticSelfImprovement3() int {
	if rand.Intn(3) < 2 {
		return 1
	}
	return 3

}
