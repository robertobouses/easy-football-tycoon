package app

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// Pass           = "Pass"
// Shot           = "Shot"
// Dribble        = "Dribble"
// Header         = "Header"
// Clearance      = "Clearance"
// Foul           = "Foul"
// FreeKick       = "Free Kick"
// PenaltyKick    = "Penalty Kick"
// ThrowIn        = "Throw In"
// CornerKick     = "Corner Kick"
// Offside        = "Offside"
// Marking        = "Marking"
// Interception   = "Interception"
// CounterAttack  = "Counter Attack"
// GoalKick       = "Goal Kick"
//YELLOW AND RED CARD
//OFFSIDE

//TODO JUGADOR RIVAL MEDIA DE STATS

//UpdatePlayerStats(playerId uuid.UUID, appearances int, blocks int, saves int, aerialduel int, keypass int, assists int, chances int, , goals int, mvp int, rating float64)

func CalculateSuccessIndividualEvent(skill int) int {
	switch {
	case skill < 8:
		return 0
	case skill >= 8 && skill < 14:
		return ProbabilisticIncrement14()
	case skill >= 14 && skill < 21:
		return ProbabilisticIncrement20()
	case skill >= 21 && skill < 30:
		return ProbabilisticIncrement25()
	case skill >= 30 && skill < 36:
		return ProbabilisticIncrement33()
	case skill >= 36 && skill < 44:
		return ProbabilisticIncrement40()
	case skill >= 44 && skill < 59:
		return ProbabilisticIncrement50()
	case skill >= 59 && skill < 68:
		return ProbabilisticIncrement66()
	case skill >= 68 && skill < 74:
		return ProbabilisticIncrement75()
	case skill >= 74 && skill < 92:
		return ProbabilisticIncrement90()

	default:
		return 1
	}
}

func CalculateSuccessConfrontation(atackerSkill, defenderSkill int) int {

	switch {
	case atackerSkill < defenderSkill-91:
		return 0
	case atackerSkill >= defenderSkill-91 && atackerSkill < defenderSkill-74:
		return ProbabilisticIncrement14()
	case atackerSkill >= defenderSkill-74 && atackerSkill < defenderSkill-69:
		return ProbabilisticIncrement20()
	case atackerSkill >= defenderSkill-69 && atackerSkill < defenderSkill-61:
		return ProbabilisticIncrement25()
	case atackerSkill >= defenderSkill-61 && atackerSkill < defenderSkill-52:
		return ProbabilisticIncrement33()
	case atackerSkill >= defenderSkill-52 && atackerSkill < defenderSkill-43:
		return ProbabilisticIncrement40()
	case atackerSkill >= defenderSkill-43 && atackerSkill < defenderSkill-30:
		return ProbabilisticIncrement44()
	case atackerSkill >= defenderSkill-30 && atackerSkill < defenderSkill-12:
		return ProbabilisticIncrement50()
	case atackerSkill >= defenderSkill-12 && atackerSkill < defenderSkill:
		return ProbabilisticIncrement57()
	case atackerSkill >= defenderSkill && atackerSkill < defenderSkill+20:
		return ProbabilisticIncrement62()
	case atackerSkill >= defenderSkill+20 && atackerSkill < defenderSkill+33:
		return ProbabilisticIncrement66()
	case atackerSkill >= defenderSkill+33 && atackerSkill < defenderSkill-37:
		return ProbabilisticIncrement71()
	case atackerSkill >= defenderSkill+37 && atackerSkill < defenderSkill+49:
		return ProbabilisticIncrement75()
	case atackerSkill >= defenderSkill+49 && atackerSkill < defenderSkill+64:
		return ProbabilisticIncrement80()
	case atackerSkill >= defenderSkill+64 && atackerSkill < defenderSkill+77:
		return ProbabilisticIncrement90()
	case atackerSkill >= defenderSkill+77 && atackerSkill < defenderSkill+96:
		return ProbabilisticIncrement94()
	case atackerSkill >= defenderSkill+96:
		return 1

	default:
		return 0
	}
}

func (a AppService) KeyPass(lineup, rivalLineup []Lineup) (string, int, int, int, int, error) {

	passer := a.GetRandomMidfielder(lineup)
	receiver := a.GetRandomForward(lineup)
	if passer == nil || receiver == nil {
		return "", 0, 0, 0, 0, fmt.Errorf("no hay suficientes jugadores disponibles para realizar un pase")
	}
	successfulPass := CalculateSuccessIndividualEvent(passer.Technique)
	var sentence string
	var lineupChances, rivalChances, lineupGoals, rivalGoals int

	if successfulPass == 1 {
		sentence := fmt.Sprintf("%s makes a key pass to %s.", passer.LastName, receiver.LastName)
		log.Println(sentence)

		err := a.statsRepo.UpdatePlayerStats(passer.PlayerId, 0, 0, 0, 0, 1, 0, 0, 0, 0, IncreaseRatingModerately)
		if err != nil {
			log.Printf("Error updating player stats for passer: %v", err)
			return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, err
		}

		err = a.statsRepo.UpdatePlayerStats(receiver.PlayerId, 0, 0, 0, 0, 0, 0, 1, 0, 0, IncreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for receiver: %v", err)
			return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, err
		}
		lineupChances = 1
		if resultOfEvent := ProbabilisticIncrement14(); resultOfEvent == 1 {
			a.PenaltyKick(lineup, rivalLineup)
		} else {
			a.Shot(lineup, rivalLineup, passer)
		}

		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
	}

	sentence = fmt.Sprintf("%s fails to make a key pass to %s.", passer.LastName, receiver.LastName)
	log.Println(sentence)

	_ = a.statsRepo.UpdatePlayerStats(passer.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, DecreaseRatingSlightly)
	_ = a.statsRepo.UpdatePlayerStats(receiver.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, DecreaseRatingSlightly)

	lineupChances = 0

	return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
}

func (a *AppService) Shot(lineup []Lineup, rivalLineup []Lineup, passer *Lineup) (string, int, int, int, int, error) {

	shooter := a.GetRandomForward(lineup)
	if shooter == nil {
		return "", 0, 0, 0, 0, errors.New("no forward player found in lineup")
	}
	goalkeeper := a.GetGoalkeeper(rivalLineup)
	if goalkeeper == nil {
		return "", 0, 0, 0, 0, errors.New("no goalkeeper found in rival lineup")
	}
	defender := a.GetRandomDefender(rivalLineup)
	if defender == nil {
		return "", 0, 0, 0, 0, errors.New("no defender player found in lineup")
	}

	log.Printf("Shooter: %+v, Defender: %+v", shooter, defender)

	successfulAgainstDefender := CalculateSuccessConfrontation(shooter.Technique, defender.Technique)

	var sentence string
	var lineupChances, rivalChances, lineupGoals, rivalGoals int

	if successfulAgainstDefender == 1 {
		err := a.statsRepo.UpdatePlayerStats(shooter.PlayerId, 0, 0, 0, 0, 0, 0, 1, 0, 0, IncreaseRatingModerately)
		if err != nil {
			log.Printf("Error updating player stats for shooter: %v", err)
			return "", 0, 0, 0, 0, err
		}
		log.Printf("%s supera a %s.\n", shooter.LastName, defender.LastName)

		successfulAgainstGoalkeeper := CalculateSuccessConfrontation(shooter.Technique, goalkeeper.Technique)

		if successfulAgainstGoalkeeper == 1 {
			sentence := fmt.Sprintf("GOOOOOL! %s scores a goal!\n", shooter.LastName)
			log.Println(sentence)

			err := a.statsRepo.UpdatePlayerStats(shooter.PlayerId, 0, 0, 0, 0, 0, 0, 0, 1, 0, IncreaseRatingDrastically)
			if err != nil {
				log.Printf("Error updating player stats for shooter: %v", err)
				return sentence, 0, 0, 0, 0, err
			}
			if passer != nil {
				err = a.statsRepo.UpdatePlayerStats(passer.PlayerId, 0, 0, 0, 0, 0, 1, 0, 0, 0, IncreaseRatingDrastically)
				if err != nil {
					log.Printf("Error updating player stats for passer: %v", err)
					return sentence, 0, 0, 0, 0, err
				}
			}
			lineupGoals = 1

		} else {
			sentence := fmt.Sprintf("%s's shot is saved by %s.\n", shooter.LastName, goalkeeper.LastName)
			log.Println(sentence)
			err := a.statsRepo.UpdatePlayerStats(goalkeeper.PlayerId, 0, 0, 1, 0, 0, 0, 0, 0, 0, IncreaseRatingConsiderably)
			if err != nil {
				log.Printf("Error updating player stats for goalkeeper: %v", err)
				return sentence, 0, 0, 0, 0, err
			}

		}
		lineupChances = 1
	} else {
		sentence := fmt.Sprintf("%s's shot is blocked by %s.\n", shooter.LastName, defender.LastName)
		log.Println(sentence)
		err := a.statsRepo.UpdatePlayerStats(defender.PlayerId, 0, 1, 0, 0, 0, 0, 0, 0, 0, IncreaseRatingModerately)
		_ = a.statsRepo.UpdatePlayerStats(shooter.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, DecreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for defender: %v", err)
			return sentence, 0, 0, 0, 0, err
		}
		lineupChances = 0
	}
	return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
}

func (a *AppService) PenaltyKick(lineup, rivalLineup []Lineup) (string, int, int, int, int, error) {
	shooter := a.GetRandomForward(lineup)
	if shooter == nil {
		return "", 0, 0, 0, 0, errors.New("no forward player found in lineup")
	}
	goalkeeper := a.GetGoalkeeper(rivalLineup)
	if goalkeeper == nil {
		return "", 0, 0, 0, 0, errors.New("no goalkeeper found in rival lineup")
	}

	increasedShooterMental := shooter.Mental + (10 * rand.Intn(3))
	decreasedGoalkeeperMental := goalkeeper.Mental - 5

	successfulPenalty := CalculateSuccessConfrontation(increasedShooterMental, decreasedGoalkeeperMental)

	var sentence string
	var lineupChances, rivalChances, lineupGoals, rivalGoals int

	if successfulPenalty == 1 {
		sentence := fmt.Sprintf("GOOOOOL! %s scores from the penalty spot!", shooter.LastName)
		log.Println(sentence)

		err := a.statsRepo.UpdatePlayerStats(shooter.PlayerId, 0, 0, 0, 0, 0, 0, 1, 1, 0, IncreaseRatingModerately)
		if err != nil {
			log.Printf("Error updating player stats for shooter: %v", err)
			return sentence, 0, 0, 0, 0, err
		}
		lineupGoals = 1
		lineupChances = 1

		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
	} else {
		sentence = fmt.Sprintf("%s's penalty is saved by %s.\n", shooter.LastName, goalkeeper.LastName)
		log.Println(sentence)

		err := a.statsRepo.UpdatePlayerStats(goalkeeper.PlayerId, 0, 0, 1, 0, 0, 0, 0, 0, 0, IncreaseRatingDrastically)
		if err != nil {
			log.Printf("Error updating player stats for goalkeeper: %v", err)
			return sentence, 0, 0, 0, 0, err
		}

		lineupChances = 1

		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
	}
}

func (a AppService) LongShot(lineup, rivalLineup []Lineup) (string, int, int, int, int, error) {

	shooter := a.GetRandomForward(lineup)
	if shooter == nil {
		return "", 0, 0, 0, 0, errors.New("no forward player found in lineup")
	}
	goalkeeper := a.GetGoalkeeper(rivalLineup)
	if goalkeeper == nil {
		return "", 0, 0, 0, 0, errors.New("no goalkeeper found in rival lineup")
	}

	decreasedShooterTechnique := shooter.Technique - (6 * rand.Intn(4))

	successfulLongShot := CalculateSuccessConfrontation(decreasedShooterTechnique, goalkeeper.Mental)

	lineupChances := 1
	rivalChances := 0
	lineupGoals := 0
	rivalGoals := 0

	if successfulLongShot == 1 {
		sentence := fmt.Sprintf("GOOOOOL! %s scores from long distance!\n", shooter.LastName)
		log.Printf("GOAL! %s scores a long shot!\n", shooter.LastName)

		err := a.statsRepo.UpdatePlayerStats(shooter.PlayerId, 0, 0, 0, 0, 0, 0, 1, 1, 0, IncreaseRatingDrastically)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", shooter.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		lineupGoals = 1
		lineupChances = 1

		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
	} else {
		sentence := fmt.Sprintf("%s's long shot is saved by %s.\n", shooter.LastName, goalkeeper.LastName)
		log.Printf("SAVE! %s's long shot is saved by %s.\n", shooter.LastName, goalkeeper.LastName)

		err := a.statsRepo.UpdatePlayerStats(goalkeeper.PlayerId, 0, 0, 1, 0, 0, 0, 0, 0, 0, IncreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for goalkeeper %s: %v", goalkeeper.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}

		lineupChances = 1

		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
	}
}

// }

// func (a AppService) LongShot(shoter, goalkeeper Player) {
// 	log.Printf("%s attempts a long shot.\n", player.LastName)
// }

// func (a AppService) Dribble(dribbler, defender Player) {
// 	log.Printf("%s performs a dribble.\n", player.LastName)
// }

// func (a AppService) AerialDuel(header, defender Player) {
// 	log.Printf("%s engages in an aerial duel.\n", player.LastName)
// }

// func (a AppService) Clearance(clearer Player) {
// 	log.Printf("%s makes a clearance.\n", player.LastName)
// }

// func (a AppService) Foul(defender Player) {
// 	log.Printf("%s commits a foul.\n", player.LastName)
// }

// func (a AppService) FreeKick(foulShooter, goalkeeper Player) {
// 	log.Printf("%s takes a free kick.\n", player.LastName)
// }

// func (a AppService) PenaltyKick(penaltyShooter, goalkeeper Player) {
// 	log.Printf("%s takes a penalty kick.\n", player.LastName)
// }

// func (a AppService) Corner(cornerShooter, receiver Player) {
// 	log.Printf("%s takes a corner in attack.\n", player.LastName)
// }

// func (a AppService) Marking(forwarder, defender Player) {
// 	log.Printf("%s is marking an opponent.\n", player.LastName)
// }

// func (a AppService) Interception(defender Player) {
// 	log.Printf("%s makes an interception.\n", player.LastName)
// }

// func (a AppService) CounterAttack(counterpuncher1, counterpuncher2, counterpuncher3 Player) {
// 	log.Printf("%s starts a counter-attack.\n", player.LastName)
// }

// func (a AppService) GoalKick(goalkeeper Player) {
// 	log.Printf("%s takes a goal kick.\n", player.LastName)}
