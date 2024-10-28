package app

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func (a *AppService) CalculateNumberOfMatchEvents() (int, error) {

	lineupStrategy, err := a.strategyRepo.GetStrategy()
	if err != nil {
		return 0, err
	}

	rivalStrategy, err := a.strategyRepo.GetStrategy()
	if err != nil {
		return 0, err
	}

	var tempoMap = map[string]int{
		"slow_tempo":     1,
		"balanced_tempo": 2,
		"fast_tempo":     3,
	}

	lineupTempo := tempoMap[lineupStrategy.GameTempo]
	rivalTempo := tempoMap[rivalStrategy.GameTempo]
	matchTempo := lineupTempo + rivalTempo
	var numberOfMatchEvents int

	switch {
	case matchTempo <= 2:
		numberOfMatchEvents = rand.Intn(6) + 3
	case matchTempo > 2 && matchTempo <= 3:
		numberOfMatchEvents = rand.Intn(8) + 4
	case matchTempo > 3 && matchTempo <= 4:
		numberOfMatchEvents = rand.Intn(9) + 6
	case matchTempo > 4 && matchTempo <= 5:
		numberOfMatchEvents = rand.Intn(9) + 9
	case matchTempo > 5 && matchTempo <= 6:
		numberOfMatchEvents = rand.Intn(11) + 12
	}

	log.Println("numberOfMatchEvents", numberOfMatchEvents)
	return numberOfMatchEvents, nil
}

func DistributeMatchEvents(lineup, rivalLineup []Lineup, numberOfMatchEvents int, homeOrAwayString string) (int, int, error) {
	log.Println("lineup en DistributeMatchEvents", lineup)
	log.Println("rival en DistributeMatchEvents", rivalLineup)
	lineupTotalQuality, err := CalculateQuality(lineup)
	if err != nil {
		return 0, 0, err
	}
	log.Println("total lineup Quality", lineupTotalQuality)
	rivalTotalQuality, err := CalculateQuality(rivalLineup)
	if err != nil {
		return 0, 0, err
	}
	log.Println("total rival Quality", rivalTotalQuality)
	allQuality := lineupTotalQuality + rivalTotalQuality
	var lineupEvents int
	lineupProportion := float64(lineupTotalQuality) / float64(allQuality)
	if homeOrAwayString == "home" {
		lineupEvents = int(lineupProportion*float64(numberOfMatchEvents)) + rand.Intn(4) + 2
	} else {
		lineupEvents = int(lineupProportion*float64(numberOfMatchEvents)) - rand.Intn(4) - 2
	}

	log.Printf("number of lineup events %v ANTES DE RANDOMFACTOR", lineupEvents)

	randomFactor := rand.Intn(11) - 5

	lineupEvents += randomFactor

	rivalEvents := numberOfMatchEvents - lineupEvents
	log.Printf("number of lineup events %v, rival events %v Despues DE RANDOMFACTOR", lineupEvents, rivalEvents)
	if lineupEvents <= 0 {
		lineupEvents = 0
	}
	if rivalEvents < 0 {
		rivalEvents = 0
	}
	log.Printf("number of lineup events %v, rival events %v", lineupEvents, rivalEvents)
	return lineupEvents, rivalEvents, nil
}

func CalculateQuality(lineup []Lineup) (int, error) {
	lineupTotalQuality := 2*(lineup[0].TotalTechnique) + 3*(lineup[0].TotalMental) + 2*(lineup[0].TotalPhysique)

	return lineupTotalQuality, nil
}

func clamp(value int, min int, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

func (a *AppService) SimulateRivalLineup(rival Rival) []Lineup {
	rivalTypeLineup := make([]Lineup, 0, 11)

	totalTechnique := 0
	totalMental := 0
	totalPhysique := 0

	goalkeeper := Lineup{
		PlayerId:  uuid.New(),
		LastName:  "goalkeeper",
		Position:  "goalkeeper",
		Technique: clamp(rival.Technique+rand.Intn(21)-20, 0, 100),
		Mental:    clamp(rival.Mental+rand.Intn(21)-20, 0, 100),
		Physique:  clamp(rival.Physique+rand.Intn(21)-20, 0, 100),
	}

	totalTechnique += goalkeeper.Technique
	totalMental += goalkeeper.Mental
	totalPhysique += goalkeeper.Physique

	rivalTypeLineup = append(rivalTypeLineup, goalkeeper)

	maxDefenders := 5
	maxMidfielders := 5
	playersAssigned := 1

	numDefenders := rand.Intn(maxDefenders-2) + 3
	if playersAssigned+numDefenders > 11 {
		numDefenders = 11 - playersAssigned
	}
	playersAssigned += numDefenders

	for i := 0; i < numDefenders; i++ {
		defender := Lineup{
			PlayerId:  uuid.New(),
			LastName:  fmt.Sprintf("defender%d", i+1),
			Position:  "defender",
			Technique: clamp(rival.Technique+rand.Intn(21)-20, 0, 100),
			Mental:    clamp(rival.Mental+rand.Intn(21)-20, 0, 100),
			Physique:  clamp(rival.Physique+rand.Intn(21)-20, 0, 100),
		}

		totalTechnique += defender.Technique
		totalMental += defender.Mental
		totalPhysique += defender.Physique

		rivalTypeLineup = append(rivalTypeLineup, defender)
	}

	numMidfielders := rand.Intn(maxMidfielders-2) + 3
	if playersAssigned+numMidfielders > 11 {
		numMidfielders = 11 - playersAssigned
	}
	playersAssigned += numMidfielders

	for i := 0; i < numMidfielders; i++ {
		midfielder := Lineup{
			PlayerId:  uuid.New(),
			LastName:  fmt.Sprintf("midfielder%d", i+1),
			Position:  "midfielder",
			Technique: clamp(rival.Technique+rand.Intn(21)-20, 0, 100),
			Mental:    clamp(rival.Mental+rand.Intn(21)-20, 0, 100),
			Physique:  clamp(rival.Physique+rand.Intn(21)-20, 0, 100),
		}

		totalTechnique += midfielder.Technique
		totalMental += midfielder.Mental
		totalPhysique += midfielder.Physique

		rivalTypeLineup = append(rivalTypeLineup, midfielder)
	}

	numForwards := 11 - playersAssigned
	for i := 0; i < numForwards; i++ {
		forward := Lineup{
			PlayerId:  uuid.New(),
			LastName:  fmt.Sprintf("forward%d", i+1),
			Position:  "forward",
			Technique: clamp(rival.Technique+rand.Intn(21)-20, 0, 100),
			Mental:    clamp(rival.Mental+rand.Intn(3), 0, 100),
			Physique:  clamp(rival.Physique+rand.Intn(3), 0, 100),
		}

		totalTechnique += forward.Technique
		totalMental += forward.Mental
		totalPhysique += forward.Physique

		rivalTypeLineup = append(rivalTypeLineup, forward)
	}

	for i := range rivalTypeLineup {
		rivalTypeLineup[i].TotalTechnique = totalTechnique
		rivalTypeLineup[i].TotalMental = totalMental
		rivalTypeLineup[i].TotalPhysique = totalPhysique
	}

	return rivalTypeLineup
}

func (a *AppService) GetRandomDefender(lineup []Lineup) *Lineup {
	defenders := filterPlayersByPosition(lineup, "defender")
	return GetRandomPlayer(defenders)
}

func (a *AppService) GetRandomMidfielder(lineup []Lineup) *Lineup {
	midfielders := filterPlayersByPosition(lineup, "midfielder")
	return GetRandomPlayer(midfielders)
}

func (a *AppService) GetRandomForward(lineup []Lineup) *Lineup {
	forwards := filterPlayersByPosition(lineup, "forward")
	return GetRandomPlayer(forwards)
}

func (a *AppService) GetGoalkeeper(lineup []Lineup) *Lineup {
	goalkeepers := filterPlayersByPosition(lineup, "goalkeeper")
	if len(goalkeepers) == 0 {
		return nil
	}
	return &goalkeepers[0]
}
func GetRandomPlayer(players []Lineup) *Lineup {
	if len(players) == 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	return &players[rand.Intn(len(players))]
}
func GetRandomPlayerExcludingGoalkeeper(lineup []Lineup) *Lineup {
	players := filterOutGoalkeepers(lineup)
	return GetRandomPlayer(players)
}

func filterOutGoalkeepers(lineup []Lineup) []Lineup {
	var filtered []Lineup
	for _, player := range lineup {
		if player.Position != "goalkeeper" {
			filtered = append(filtered, player)
		}
	}
	return filtered
}

type Event struct {
	Name    string
	Execute func() (string, int, int, int, int, error)
}

type EventResult struct {
	Event  string `json:"event"`
	Minute int    `json:"minute"`
	Team   string `json:"team"`
}

func (a AppService) GenerateEvents(lineup, rivalLineup []Lineup, numberOfLineupEvents, numberOfRivalEvents int, rivalName string) ([]EventResult, []EventResult, int, int, int, int) {

	lineupEvents := []Event{
		{
			"Pase clave",
			func() (string, int, int, int, int, error) {
				return a.KeyPass(lineup, rivalLineup)
			},
		},
		{
			"Remate a puerta",
			func() (string, int, int, int, int, error) {
				return a.Shot(lineup, rivalLineup, a.GetRandomForward(lineup))
			},
		},
		{
			"Penalty",
			func() (string, int, int, int, int, error) {
				return a.PenaltyKick(lineup, rivalLineup)
			},
		},
		{
			"Tiro lejano",
			func() (string, int, int, int, int, error) {
				return a.LongShot(lineup, rivalLineup)
			},
		},
		{
			"Falta Indirecta",
			func() (string, int, int, int, int, error) {
				return a.IndirectFreeKick(lineup, rivalLineup)
			},
		},
	}

	rivalEvents := []Event{
		{
			"Pase clave",
			func() (string, int, int, int, int, error) {
				return a.KeyPass(rivalLineup, lineup)
			},
		},
		{
			"Remate a puerta",
			func() (string, int, int, int, int, error) {
				return a.Shot(rivalLineup, lineup, a.GetRandomForward(rivalLineup))
			},
		},
		{
			"Penalty",
			func() (string, int, int, int, int, error) {
				return a.PenaltyKick(rivalLineup, lineup)
			},
		},
		{
			"Tiro lejano",
			func() (string, int, int, int, int, error) {
				return a.LongShot(rivalLineup, lineup)
			},
		},
		{
			"Falta Indirecta",
			func() (string, int, int, int, int, error) {
				return a.IndirectFreeKick(rivalLineup, lineup)
			},
		},
	}
	var lineupResults []EventResult
	var rivalResults []EventResult
	var lineupChances, rivalChances, lineupGoals, rivalGoals int

	for i := 0; i < numberOfLineupEvents; i++ {
		event := lineupEvents[rand.Intn(len(lineupEvents))]
		result, newLineupChances, newRivalChances, newLineupGoals, newRivalGoals, err := event.Execute()
		if err != nil {
			fmt.Printf("Error executing lineup event: %v\n", err)
			continue
		}
		if result == "" {
			fmt.Println("Generated empty event for lineup!")
		} else {
			fmt.Printf("Generated lineup event: %s\n", result)
		}
		lineupChances += newLineupChances
		rivalChances += newRivalChances
		lineupGoals += newLineupGoals
		rivalGoals += newRivalGoals

		minute := rand.Intn(90)
		lineupResults = append(lineupResults, EventResult{
			Event:  result + "para tu equipo",
			Minute: minute,
			Team:   "tu equipo",
		})
		fmt.Printf("Generated event: %s at minute %d\n", result, minute)

	}
	for i := 0; i < numberOfRivalEvents; i++ {
		event := rivalEvents[rand.Intn(len(rivalEvents))]
		result, newRivalChances, newLineupChances, newRivalGoals, newLineupGoals, err := event.Execute()
		if err != nil {
			fmt.Printf("Error executing rival event: %v\n", err)
			continue
		}

		lineupChances += newLineupChances
		rivalChances += newRivalChances
		lineupGoals += newLineupGoals
		rivalGoals += newRivalGoals

		minute := rand.Intn(90)
		rivalResults = append(rivalResults, EventResult{
			Event:  result + " para " + rivalName,
			Minute: minute,
			Team:   rivalName,
		})
		fmt.Printf("Generated event: %s at minute %d\n", result, minute)

	}

	return lineupResults, rivalResults, lineupChances, rivalChances, lineupGoals, rivalGoals
}
