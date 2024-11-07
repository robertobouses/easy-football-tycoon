package app

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

// Header         = "Header"
// Clearance      = "Clearance"
// Marking        = "Marking"
// Interception   = "Interception"
// CounterAttack  = "Counter Attack"
// GoalKick       = "Goal Kick"

//UpdatePlayerStats(playerId uuid.UUID, appearances int, blocks int, saves int, aerialduel int, keypass int, assists int, chances int, , goals int, mvp int, rating float64)

func CalculateSuccessIndividualEvent(skill int) int {
	log.Printf("Evaluating success of individual event for skill level: %d", skill)

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
	log.Printf("Calculating confrontation success: Attacker skill = %d, Defender skill = %d", atackerSkill, defenderSkill)

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
	log.Printf("Selected passer: %+v, receiver: %+v", passer, receiver)

	if passer == nil || receiver == nil {
		return "no hay suficientes jugadores disponibles para realizar un pase", 0, 0, 0, 0, fmt.Errorf("no hay suficientes jugadores disponibles para realizar un pase")
	}
	successfulPass := CalculateSuccessIndividualEvent(passer.Technique)
	var sentence string
	var lineupChances, rivalChances, lineupGoals, rivalGoals int

	if successfulPass == 1 {
		log.Printf("Pass success calculated: %d", successfulPass)
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
	//TODO FUNCION QUE IMPRIME UN VACÍO
	shooter := a.GetRandomForward(lineup)
	if shooter == nil {
		return "no forward player found in lineup", 0, 0, 0, 0, errors.New("no forward player found in lineup")
	}
	goalkeeper := a.GetGoalkeeper(rivalLineup)
	if goalkeeper == nil {
		return "no goalkeeper found in rival lineup", 0, 0, 0, 0, errors.New("no goalkeeper found in rival lineup")
	}
	defender := a.GetRandomDefender(rivalLineup)
	if defender == nil {
		return "no defender player found in lineup", 0, 0, 0, 0, errors.New("no defender player found in lineup")
	}

	log.Printf("Shooter: %+v, Defender: %+v, Goalkeeper: %+v", shooter, defender, goalkeeper)

	var sentence string
	var lineupChances, rivalChances, lineupGoals, rivalGoals int

	successfulAgainstDefender := CalculateSuccessConfrontation(shooter.Technique, defender.Technique)
	log.Printf("Success against defender: %d", successfulAgainstDefender)

	if successfulAgainstDefender == 1 {
		sentence = fmt.Sprintf("%s escapes from %s's marking...", shooter.LastName, defender.LastName)
		lineupChances = 1
		log.Println("el delantero supera al defensa")
		err := a.statsRepo.UpdatePlayerStats(shooter.PlayerId, 0, 0, 0, 0, 0, 0, 1, 0, 0, IncreaseRatingModerately)
		if err != nil {
			log.Printf("Error updating player stats for shooter: %v", err)
			return "Error updating player stats for shooter", 0, 0, 0, 0, err
		}
		log.Printf("%s supera a %s.\n", shooter.LastName, defender.LastName)

		successfulAgainstGoalkeeper := CalculateSuccessConfrontation(shooter.Technique, goalkeeper.Technique)

		if successfulAgainstGoalkeeper == 1 {
			sentence += fmt.Sprintf(" %s shoots and also beats the goalkeeper... GOOOOOL! %s is just a spectator in the play %s scores a goal!\n", shooter.LastName, goalkeeper.LastName, shooter.LastName)
			log.Println(sentence)
			log.Println("el delantero también supera al portero, es GOOL")
			lineupGoals = 1

			err := a.statsRepo.UpdatePlayerStats(shooter.PlayerId, 0, 0, 0, 0, 0, 0, 0, 1, 0, IncreaseRatingDrastically)
			if err != nil {
				log.Printf("Error updating player stats for shooter: %v", err)
				return sentence, 0, 0, 0, 0, err
			}
			if passer != nil {
				log.Printf("El pasador existe y es: %v", passer)
				err = a.statsRepo.UpdatePlayerStats(passer.PlayerId, 0, 0, 0, 0, 0, 1, 0, 0, 0, IncreaseRatingDrastically)
				if err != nil {
					log.Printf("Error updating player stats for passer: %v", err)
					return sentence, 0, 0, 0, 0, err
				}
			}

			return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
		} else {
			sentence += fmt.Sprintf(" %s's shot is saved by %s.\n", shooter.LastName, goalkeeper.LastName)
			log.Println(sentence)
			err := a.statsRepo.UpdatePlayerStats(goalkeeper.PlayerId, 0, 0, 1, 0, 0, 0, 0, 0, 0, IncreaseRatingConsiderably)
			if err != nil {
				log.Printf("Error updating player stats for goalkeeper: %v", err)
				return sentence, 0, 0, 0, 0, err
			}

		}
		lineupChances = 1
	} else {
		log.Printf("el defensa bloqueó el disparo en primera instancia")
		sentence += fmt.Sprintf(" %s's shot is blocked by %s.\n", shooter.LastName, defender.LastName)
		log.Println(sentence)
		err := a.statsRepo.UpdatePlayerStats(defender.PlayerId, 0, 1, 0, 0, 0, 0, 0, 0, 0, IncreaseRatingModerately)
		_ = a.statsRepo.UpdatePlayerStats(shooter.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, DecreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for defender: %v", err)
			return sentence, 0, 0, 0, 0, err
		}
		lineupChances = 0
		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
	}
	return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
}

func (a *AppService) PenaltyKick(lineup, rivalLineup []Lineup) (string, int, int, int, int, error) {
	shooter := a.GetRandomForward(lineup)
	if shooter == nil {
		return "no forward player found in lineup", 0, 0, 0, 0, errors.New("no forward player found in lineup")
	}
	goalkeeper := a.GetGoalkeeper(rivalLineup)
	if goalkeeper == nil {
		return "no goalkeeper found in rival lineup", 0, 0, 0, 0, errors.New("no goalkeeper found in rival lineup")
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
		return "no forward player found in lineup", 0, 0, 0, 0, errors.New("no forward player found in lineup")
	}
	goalkeeper := a.GetGoalkeeper(rivalLineup)
	if goalkeeper == nil {
		return "no goalkeeper found in rival lineup", 0, 0, 0, 0, errors.New("no goalkeeper found in rival lineup")
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

func (a AppService) IndirectFreeKick(lineup, rivalLineup []Lineup) (string, int, int, int, int, error) {

	shooter := a.GetRandomMidfielder(lineup)
	if shooter == nil {
		return "no shooter player found in lineup", 0, 0, 0, 0, errors.New("no shooter player found in lineup")
	}
	defenderOnAttack := a.GetRandomDefender(lineup)
	if defenderOnAttack == nil {
		return "no defender player found in lineup", 0, 0, 0, 0, errors.New("no defender player found in lineup")
	}
	rivalDefender := a.GetRandomDefender(rivalLineup)
	if rivalDefender == nil {
		return "no rivalDefender player found in lineup", 0, 0, 0, 0, errors.New("no rivalDefender player found in lineup")
	}
	goalkeeper := a.GetGoalkeeper(rivalLineup)
	if goalkeeper == nil {
		return "no goalkeeper found in rival lineup", 0, 0, 0, 0, errors.New("no goalkeeper found in rival lineup")
	}

	increasedShooterTechnique := shooter.Technique + (4 * rand.Intn(6))
	increasedRivalDefenderPhysique := rivalDefender.Physique + rand.Intn(30)

	attackAtributes := increasedShooterTechnique + defenderOnAttack.Physique
	defenseAtributes := increasedRivalDefenderPhysique + goalkeeper.Technique

	successfulLongShot := CalculateSuccessConfrontation(attackAtributes, defenseAtributes)

	lineupChances := 1
	rivalChances := 0
	lineupGoals := 0
	rivalGoals := 0

	if successfulLongShot == 1 {
		sentence := fmt.Sprintf("%s takes the free kick... It is a center to the area... %s and %s jump to fight for the center... %s head the ball...%s can't do anything...  GOOOOOL! %s scores!\n", shooter.LastName, defenderOnAttack.LastName, rivalDefender.LastName, defenderOnAttack.LastName, goalkeeper.LastName, defenderOnAttack.LastName)
		log.Printf("GOAL! %s scores a long shot!\n", shooter.LastName)

		err := a.statsRepo.UpdatePlayerStats(shooter.PlayerId, 0, 0, 0, 0, 1, 1, 0, 0, 0, IncreaseRatingConsiderably)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", shooter.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		err = a.statsRepo.UpdatePlayerStats(defenderOnAttack.PlayerId, 0, 0, 0, 1, 0, 0, 1, 1, 0, IncreaseRatingDrastically)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", shooter.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		err = a.statsRepo.UpdatePlayerStats(rivalDefender.PlayerId, 0, 0, 0, 1, 0, 0, 0, 0, 0, DecreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", shooter.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		err = a.statsRepo.UpdatePlayerStats(goalkeeper.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, DecreaseRatingModerately)
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

		err := a.statsRepo.UpdatePlayerStats(shooter.PlayerId, 0, 0, 0, 0, 1, 0, 0, 0, 0, DecreaseRatingModerately)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", shooter.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		err = a.statsRepo.UpdatePlayerStats(defenderOnAttack.PlayerId, 0, 0, 0, 1, 0, 0, 1, 0, 0, DecreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", shooter.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		err = a.statsRepo.UpdatePlayerStats(rivalDefender.PlayerId, 0, 0, 0, 1, 0, 0, 0, 0, 0, IncreaseRatingModerately)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", shooter.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		err = a.statsRepo.UpdatePlayerStats(goalkeeper.PlayerId, 0, 0, 1, 0, 0, 0, 0, 0, 0, IncreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", shooter.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}

		lineupChances = 1

		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
	}
}

func (a AppService) Dribble(lineup, rivalLineup []Lineup) (string, int, int, int, int, error) {
	var dribbler, defender *Lineup
	var sentence string
	var lineupChances, rivalChances, lineupGoals, rivalGoals int

	prob := ProbabilisticIncrement20() + ProbabilisticIncrement20()
	if prob <= 0 {
		dribbler = a.GetRandomMidfielder(lineup)
		defender = a.GetRandomMidfielder(rivalLineup)
	} else if prob <= 1 {
		dribbler = a.GetRandomForward(lineup)
		defender = a.GetRandomDefender(rivalLineup)
	} else if prob > 1 {
		dribbler = a.GetRandomMidfielder(lineup)
		defender = a.GetRandomDefender(rivalLineup)
	}

	log.Printf("Selected passer: %+v, receiver: %+v", dribbler, defender)

	if dribbler == nil || defender == nil {
		return "no hay suficientes jugadores disponibles para realizar un pase", 0, 0, 0, 0, fmt.Errorf("no hay suficientes jugadores disponibles para realizar un pase")
	}

	sentence = fmt.Sprintf("%s tries a dribbling", dribbler.LastName)
	successfulDribble := CalculateSuccessIndividualEvent(dribbler.Technique)

	if successfulDribble == 1 {
		log.Printf("Pass success calculated: %d", successfulDribble)
		sentence += " and succeeds..."
		log.Println(sentence)

		err := a.statsRepo.UpdatePlayerStats(dribbler.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, IncreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for dribbler: %v", err)
			return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, err
		}
		successfulConfrontation := CalculateSuccessConfrontation(dribbler.Technique, defender.Technique)
		if successfulConfrontation == 1 {
			sentence += fmt.Sprintf(" %s dribbled %s...", dribbler.LastName, defender.LastName)

			err := a.statsRepo.UpdatePlayerStats(dribbler.PlayerId, 0, 0, 0, 0, 0, 0, 1, 0, 0, IncreaseRatingSlightly)
			if err != nil {
				log.Printf("Error updating player stats for dribbler %s: %v", dribbler.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
			err = a.statsRepo.UpdatePlayerStats(defender.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, DecreaseRatingSlightly)
			if err != nil {
				log.Printf("Error updating player stats for dribbler %s: %v", dribbler.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
			lineupChances = 1

			if resultOfEvent := ProbabilisticIncrement40(); resultOfEvent == 1 {
				sentence += " the occasion ends with a shot"
				log.Println("the occasion ends with a shot")
				a.Shot(lineup, rivalLineup, dribbler)
			} else {
				sentence += " the occasion ends with a foul"
				log.Println("the occasion ends with a shot")
				a.currentePlayerOnMatchEvent = defender
				a.Foul(lineup, rivalLineup, defender)
			}

			return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
		} else {
			sentence += fmt.Sprintf(" but %s lost the dribbled against %s", dribbler.LastName, defender.LastName)

			err := a.statsRepo.UpdatePlayerStats(dribbler.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, DecreaseRatingModerately)
			if err != nil {
				log.Printf("Error updating player stats for dribbler %s: %v", dribbler.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
			err = a.statsRepo.UpdatePlayerStats(defender.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, IncreaseRatingModerately)
			if err != nil {
				log.Printf("Error updating player stats for dribbler %s: %v", dribbler.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
		}

		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
	}
	sentence += fmt.Sprintf(" Oh Noo, %s trips over the ball, and fails the dribble", dribbler.LastName)

	log.Println(sentence)

	_ = a.statsRepo.UpdatePlayerStats(dribbler.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, DecreaseRatingSlightly)
	_ = a.statsRepo.UpdatePlayerStats(defender.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, DecreaseRatingSlightly)

	lineupChances = 0

	return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
}

func (a AppService) Foul(lineup, rivalLineup []Lineup, defender *Lineup) (string, int, int, int, int, error) {
	var sentence string
	var lineupChances, rivalChances, lineupGoals, rivalGoals int

	probabilyYellowOrRedCard := ProbabilisticIncrement40() + ProbabilisticIncrement20()

	resultOfEvent := ProbabilisticIncrement50() + ProbabilisticIncrement33() + ProbabilisticIncrement33()

	if probabilyYellowOrRedCard >= 1 {
		sentence = "the referee puts his hand in his pocket"
		a.currentePlayerOnMatchEvent = defender
		a.YellowOrRedCard(lineup, defender)
	}
	if resultOfEvent >= 2 {
		sentence = "the foul is in the middle of the field"
		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
	}
	if resultOfEvent >= 1 {
		sentence = "the foul is in the middle of the field"
		a.IndirectFreeKick(lineup, rivalLineup)

	} else {
		sentence = "the foul is in a dangerous area of the field"
		a.DirectFreeKick(lineup, rivalLineup)
	}
	return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil

}

func (a AppService) YellowOrRedCard(lineup []Lineup, defender *Lineup) (string, int, int, int, int, error) {
	var sentence string
	var lineupChances, rivalChances, lineupGoals, rivalGoals, probabilyYellowCard int
	sentence = "The referee puts his hand in his pocket"

	if defender == nil {
		defender = a.GetRandomDefender(lineup)
		if defender == nil {
			return "no defender player found in lineup", 0, 0, 0, 0, errors.New("no defender player found in lineup")
		}
	}

	probabilyIncrementByAgressive := CalculateSuccessIndividualEvent(defender.Mental)

	if probabilyIncrementByAgressive >= 1 {
		probabilyYellowCard = ProbabilisticIncrement62()
	} else {
		probabilyYellowCard = ProbabilisticIncrement75()
	}

	if probabilyYellowCard >= 1 {
		sentence += fmt.Sprintf("The referee gives %v a yellow card", defender.LastName)
		err := a.statsRepo.UpdatePlayerStats(defender.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, DecreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for shooter: %v", err)
			return sentence, 0, 0, 0, 0, err
		}

	} else {

		sentence += fmt.Sprintf("The referee gives %v a red card", defender.LastName)
		log.Println("el jugador fue expulsado de la alineación")
		err := a.lineupRepo.DeletePlayerFromLineup(defender.PlayerId)
		if err != nil {
			log.Printf("Error delete player: %v", err)
			return sentence, 0, 0, 0, 0, err
		}
		err = a.statsRepo.UpdatePlayerStats(defender.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, DecreaseRatingDrastically)
		if err != nil {
			log.Printf("Error updating player stats: %v", err)
			return sentence, 0, 0, 0, 0, err
		}
	}

	return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil

}

func (a AppService) DirectFreeKick(lineup, rivalLineup []Lineup) (string, int, int, int, int, error) {

	shooter := a.GetRandomForward(lineup)
	if shooter == nil {
		return "no forward player found in lineup", 0, 0, 0, 0, errors.New("no forward player found in lineup")
	}
	goalkeeper := a.GetGoalkeeper(rivalLineup)
	if goalkeeper == nil {
		return "no goalkeeper found in rival lineup", 0, 0, 0, 0, errors.New("no goalkeeper found in rival lineup")
	}

	decreasedShooterTechnique := shooter.Technique - (6 * rand.Intn(7))

	successfulLongShot := CalculateSuccessConfrontation(decreasedShooterTechnique, goalkeeper.Technique)

	lineupChances := 1
	rivalChances := 0
	lineupGoals := 0
	rivalGoals := 0

	if successfulLongShot == 1 {
		sentence := fmt.Sprintf("GOOOOOL! %s scores from direct free kick distance!\n", shooter.LastName)
		log.Printf("GOAL! %s scores a free kick long shot!\n", shooter.LastName)

		err := a.statsRepo.UpdatePlayerStats(shooter.PlayerId, 0, 0, 0, 0, 0, 0, 1, 1, 0, IncreaseRatingDrastically+IncreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", shooter.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		lineupGoals = 1
		lineupChances = 1

		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
	} else {
		sentence := fmt.Sprintf("%s's free kick shot is saved by %s.\n", shooter.LastName, goalkeeper.LastName)
		log.Printf("SAVE! %s's free kick is saved by %s.\n", shooter.LastName, goalkeeper.LastName)

		err := a.statsRepo.UpdatePlayerStats(goalkeeper.PlayerId, 0, 0, 1, 0, 0, 0, 0, 0, 0, IncreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for goalkeeper %s: %v", goalkeeper.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}

		lineupChances = 1

		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
	}
}

func (a AppService) GreatScoringChance(lineup []Lineup) (string, int, int, int, int, error) {
	var shooter *Lineup
	var sentence string
	var lineupChances, rivalChances, lineupGoals, rivalGoals int
	prob := ProbabilisticIncrement66()
	if prob == 1 {
		shooter = a.GetRandomForward(lineup)
	} else {
		shooter = a.GetRandomMidfielder(lineup)
	}
	prob = ProbabilisticIncrement71()
	lineupChances = 1
	if prob == 1 {
		lineupGoals = 1
		sentence = fmt.Sprintf("%s score a great easy chance", shooter.LastName)
		err := a.statsRepo.UpdatePlayerStats(shooter.PlayerId, 0, 0, 0, 0, 0, 0, 1, 1, 0, IncreaseRatingModerately)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", shooter.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
	} else {
		sentence = fmt.Sprintf("%s fails miserably with a very clear scoring chance", shooter.LastName)
		err := a.statsRepo.UpdatePlayerStats(shooter.PlayerId, 0, 0, 0, 0, 0, 0, 1, 0, 0, DecreaseRatingConsiderably)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", shooter.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil

	}
}

func (a AppService) CornerKick(lineup, rivalLineup []Lineup) (string, int, int, int, int, error) {
	var centerer *Lineup
	var attacker, defender *Lineup
	var sentence string
	var lineupChances, rivalChances, lineupGoals, rivalGoals int

	centerer = a.GetRandomMidfielder(lineup)

	incrementedTechnique := centerer.Technique + rand.Intn(20)
	prob := CalculateSuccessIndividualEvent(incrementedTechnique)
	lineupChances = 1

	if prob == 1 {
		defender = a.GetRandomDefender(rivalLineup)
		prob = ProbabilisticIncrement62()
		if prob == 1 {
			attacker = a.GetRandomDefender(lineup)
		} else {
			attacker = a.GetRandomMidfielder(lineup)
		}
		prob = CalculateSuccessConfrontation(attacker.Physique, defender.Physique)
		if prob == 1 {
			sentence = fmt.Sprintf("GOOOOOL, %s took the corner very well, and %s beats %s with a incredible jump and heads at goal", centerer.LastName, attacker.LastName, defender.LastName)
			lineupGoals = 1
			err := a.statsRepo.UpdatePlayerStats(centerer.PlayerId, 0, 0, 0, 0, 1, 1, 0, 0, 0, IncreaseRatingModerately)
			if err != nil {
				log.Printf("Error updating player stats for shooter %s: %v", centerer.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
			err = a.statsRepo.UpdatePlayerStats(attacker.PlayerId, 0, 0, 0, 1, 0, 0, 1, 1, 0, IncreaseRatingConsiderably)
			if err != nil {
				log.Printf("Error updating player stats for shooter %s: %v", centerer.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
			err = a.statsRepo.UpdatePlayerStats(defender.PlayerId, 0, 0, 0, 1, 0, 0, 0, 0, 0, DecreaseRatingModerately)
			if err != nil {
				log.Printf("Error updating player stats for shooter %s: %v", centerer.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
			return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
		} else {
			sentence = fmt.Sprintf("%s takes the corner... but %s beats %s to the jump and clears the ball", centerer.LastName, defender.LastName, attacker.LastName)
			err := a.statsRepo.UpdatePlayerStats(centerer.PlayerId, 0, 0, 0, 0, 1, 0, 0, 0, 0, IncreaseRatingSlightly)
			if err != nil {
				log.Printf("Error updating player stats for shooter %s: %v", centerer.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
			err = a.statsRepo.UpdatePlayerStats(attacker.PlayerId, 0, 0, 0, 1, 0, 0, 0, 0, 0, DecreaseRatingSlightly)
			if err != nil {
				log.Printf("Error updating player stats for shooter %s: %v", centerer.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
			err = a.statsRepo.UpdatePlayerStats(defender.PlayerId, 0, 0, 0, 1, 0, 0, 0, 0, 0, IncreaseRatingModerately)
			if err != nil {
				log.Printf("Error updating player stats for shooter %s: %v", centerer.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
			return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
		}
	} else {
		sentence = fmt.Sprintf("the corner was wasted by %s", centerer.LastName)
		err := a.statsRepo.UpdatePlayerStats(centerer.PlayerId, 0, 0, 0, 0, 1, 0, 0, 0, 0, DecreaseRatingModerately)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", centerer.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		return sentence, lineupChances, rivalChances, lineupGoals, rivalGoals, nil
	}
}

func (a AppService) InjuryDuringMatch(lineup []Lineup) (string, int, int, int, int, error) {
	var injuredPlayer *Lineup
	var sentence string
	injuredPlayer = a.GetRandomPlayerExcludingGoalkeeper(lineup)

	player, err := a.ProcessInjury(injuredPlayer.PlayerId)
	if err != nil {
		log.Printf("error al procesar la lesión, valor para error: %v, valor para jugador: %v", err, player)
	}
	sentence = fmt.Sprintf("There is a player lying on the ground... wow he is %s, he looks like he will need assistance...", injuredPlayer.LastName)
	err = a.lineupRepo.DeletePlayerFromLineup(injuredPlayer.PlayerId)
	if err != nil {
		log.Println("error al borrar el jugador de la alineación durante la lesión en partido", err)
	}
	return sentence, 0, 0, 0, 0, nil
}

func (a AppService) Offside(lineup, rivalLineup []Lineup) (string, int, int, int, int, error) {
	var passer, playerOffside *Lineup
	var lineupChances int
	var sentence string

	passer = a.GetRandomMidfielder(lineup)
	playerOffside = a.GetRandomForward(lineup)

	lineupChances = 1

	sentence = fmt.Sprintf("%s looks a pass... ", passer.LastName)
	sentence += fmt.Sprintf("%s runs behind the rival defense", playerOffside.LastName)

	prob := ProbabilisticIncrement66()
	if prob >= 1 {
		sentence += "%s its offside, the opportunity is lost"
		err := a.statsRepo.UpdatePlayerStats(passer.PlayerId, 0, 0, 0, 0, 1, 0, 0, 0, 0, DecreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", passer.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		err = a.statsRepo.UpdatePlayerStats(playerOffside.PlayerId, 0, 0, 0, 0, 0, 0, 0, 0, 0, DecreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", playerOffside.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
	} else {
		sentence += "great pass bordering on offside"
		err := a.statsRepo.UpdatePlayerStats(passer.PlayerId, 0, 0, 0, 0, 1, 0, 0, 0, 0, IncreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", passer.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		err = a.statsRepo.UpdatePlayerStats(playerOffside.PlayerId, 0, 0, 0, 0, 0, 0, 1, 0, 0, IncreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for shooter %s: %v", playerOffside.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		a.Shot(lineup, rivalLineup, passer)
	}

	return sentence, lineupChances, 0, 0, 0, nil
}

func (a AppService) Headed(lineup, rivalLineup []Lineup) (string, int, int, int, int, error) {
	var header, rivalHeader *Lineup
	var sentence string
	var lineupChances, rivalChances int

	header = a.GetRandomPlayerExcludingGoalkeeper(lineup)
	rivalHeader = a.GetRandomPlayerExcludingGoalkeeper(rivalLineup)

	sentence = "The ball comes through the air, here we have an aerial duel"

	success := CalculateSuccessConfrontation(header.Physique, rivalHeader.Physique)
	if success == 1 {
		lineupChances = 1
		sentence += fmt.Sprintf("%s wins a header in midfield against %s", header.LastName, rivalHeader.LastName)
		err := a.statsRepo.UpdatePlayerStats(header.PlayerId, 0, 0, 0, 1, 1, 0, 0, 0, 0, IncreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for header %s: %v", header.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		err = a.statsRepo.UpdatePlayerStats(rivalHeader.PlayerId, 0, 0, 0, 1, 0, 0, 0, 0, 0, DecreaseRatingSlightly)
		if err != nil {
			log.Printf("Error updating player stats for rivalHeader %s: %v", rivalHeader.LastName, err)
			return sentence, 0, 0, 0, 0, err
		}
		a.LongShot(lineup, rivalLineup)
	} else {
		sentence += fmt.Sprintf("%s loses a header in midfield against %s", header.LastName, rivalHeader.LastName)
		prob := ProbabilisticIncrement75()
		if prob >= 1 {
			rivalChances = 1
			sentence += fmt.Sprintf("%s makes a long pass, and his teammates run away", rivalHeader.LastName)
			err := a.statsRepo.UpdatePlayerStats(header.PlayerId, 0, 0, 0, 1, 0, 0, 0, 0, 0, DecreaseRatingSlightly)
			if err != nil {
				log.Printf("Error updating player stats for header %s: %v", header.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
			err = a.statsRepo.UpdatePlayerStats(rivalHeader.PlayerId, 0, 0, 0, 1, 1, 0, 0, 0, 0, IncreaseRatingSlightly)
			if err != nil {
				log.Printf("Error updating player stats for rivalHeader %s: %v", rivalHeader.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
			a.CounterAttack(rivalLineup, lineup)

		} else {
			sentence += fmt.Sprintf("%s kick the ball into the air, and there are no second plays", rivalHeader.LastName)
			err := a.statsRepo.UpdatePlayerStats(header.PlayerId, 0, 0, 0, 1, 0, 0, 0, 0, 0, IncreaseRatingSlightly)
			if err != nil {
				log.Printf("Error updating player stats for header %s: %v", header.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
			err = a.statsRepo.UpdatePlayerStats(rivalHeader.PlayerId, 0, 0, 0, 1, 0, 0, 0, 0, 0, DecreaseRatingSlightly)
			if err != nil {
				log.Printf("Error updating player stats for rivalHeader %s: %v", rivalHeader.LastName, err)
				return sentence, 0, 0, 0, 0, err
			}
		}

	}
	return sentence, lineupChances, rivalChances, 0, 0, nil
}

func (a AppService) CounterAttack(lineup, rivalLineup []Lineup) (string, int, int, int, int, error) {
	var sentence string

	sentence = "Some players run out in counterattack"
	prob := ProbabilisticIncrement66()
	if prob >= 1 {
		a.LongShot(lineup, rivalLineup)
	} else {
		prob := ProbabilisticIncrement57()
		if prob >= 1 {
			sentence += "The rival stopped the counterattack with a foul"
			a.IndirectFreeKick(lineup, rivalLineup)
		} else {
			sentence += "The opponent breaks the counterattack cleanly"
		}
	}

	return sentence, 0, 0, 0, 0, nil
}
