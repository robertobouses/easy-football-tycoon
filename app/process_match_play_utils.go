package app

import (
	"fmt"
	"math/rand"
	"sort"
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

	return numberOfMatchEvents, nil
}

func DistributeMatchEvents(lineup, rival []Lineup, numberOfMatchEvents int, homeOrAwayString string) (int, int, error) {
	lineupTotalQuality, _, allQuality, err := CalculateQuality(lineup, rival)
	if err != nil {
		return 0, 0, err
	}
	var lineupEvents int
	lineupProportion := float64(lineupTotalQuality) / float64(allQuality)
	if homeOrAwayString == "home" {
		lineupEvents = int(lineupProportion*float64(numberOfMatchEvents)) + rand.Intn(4) + 2
	} else {
		lineupEvents = int(lineupProportion*float64(numberOfMatchEvents)) - rand.Intn(4) - 2
	}

	randomFactor := rand.Intn(11) - 5

	lineupEvents += randomFactor

	rivalEvents := numberOfMatchEvents - lineupEvents

	if lineupEvents <= 0 {
		lineupEvents = 0
	}
	if rivalEvents < 0 {
		rivalEvents = 0
	}
	return lineupEvents, rivalEvents, nil
}

func CalculateQuality(lineup, rival []Lineup) (int, int, int, error) {
	lineupTotalQuality := 2*(lineup[0].TotalTechnique) + 3*(lineup[0].TotalMental) + 2*(lineup[0].TotalPhysique)
	rivalTotalQuality := 2*(rival[0].TotalTechnique) + 3*(rival[0].TotalMental) + 2*(rival[0].TotalPhysique)
	allQuality := lineupTotalQuality + rivalTotalQuality
	return lineupTotalQuality, rivalTotalQuality, allQuality, nil
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

	rivalTypeLineup = append(rivalTypeLineup, Lineup{
		PlayerId:  uuid.New(),
		LastName:  "Goalkeeper",
		Position:  "Goalkeeper",
		Technique: clamp(rival.Technique+rand.Intn(21)-20, 0, 100),
		Mental:    clamp(rival.Mental+rand.Intn(21)-20, 0, 100),
		Physique:  clamp(rival.Physique+rand.Intn(21)-20, 0, 100),
	})

	maxDefenders := 5
	maxMidfielders := 5

	playersAssigned := 1

	numDefenders := rand.Intn(maxDefenders-2) + 3
	if playersAssigned+numDefenders > 11 {
		numDefenders = 11 - playersAssigned
	}
	playersAssigned += numDefenders

	for i := 0; i < numDefenders; i++ {
		rivalTypeLineup = append(rivalTypeLineup, Lineup{
			PlayerId:  uuid.New(),
			LastName:  "Defender" + string(i+1),
			Position:  "Defender",
			Technique: clamp(rival.Technique+rand.Intn(21)-20, 0, 100),
			Mental:    clamp(rival.Mental+rand.Intn(21)-20, 0, 100),
			Physique:  clamp(rival.Physique+rand.Intn(21)-20, 0, 100),
		})
	}

	numMidfielders := rand.Intn(maxMidfielders-2) + 3
	if playersAssigned+numMidfielders > 11 {
		numMidfielders = 11 - playersAssigned
	}
	playersAssigned += numMidfielders

	for i := 0; i < numMidfielders; i++ {
		rivalTypeLineup = append(rivalTypeLineup, Lineup{
			PlayerId:  uuid.New(),
			LastName:  "Midfielder" + string(i+1),
			Position:  "Midfielder",
			Technique: clamp(rival.Technique+rand.Intn(21)-20, 0, 100),
			Mental:    clamp(rival.Mental+rand.Intn(21)-20, 0, 100),
			Physique:  clamp(rival.Physique+rand.Intn(21)-20, 0, 100),
		})
	}

	numForwards := 11 - playersAssigned
	for i := 0; i < numForwards; i++ {
		rivalTypeLineup = append(rivalTypeLineup, Lineup{
			PlayerId:  uuid.New(),
			LastName:  "Forward" + string(i+1),
			Position:  "Forward",
			Technique: clamp(rival.Technique+rand.Intn(21)-20, 0, 100),
			Mental:    clamp(rival.Mental+rand.Intn(3), 0, 100),
			Physique:  clamp(rival.Physique+rand.Intn(3), 0, 100),
		})
	}

	return rivalTypeLineup
}

type Event struct {
	Minute      int
	Description string
}

func (a *AppService) SimulateMatchClock(numEvents int) []Event {
	rand.Seed(time.Now().UnixNano())
	events := make([]Event, 0, numEvents)

	for i := 0; i < numEvents; i++ {
		minute := rand.Intn(90) + 1
		event := Event{
			Minute:      minute,
			Description: fmt.Sprintf("Evento en el minuto %d", minute),
		}
		events = append(events, event)
	}

	events = append(events, Event{Minute: 45, Description: "Descanso"})
	events = append(events, Event{Minute: 90, Description: "Final del partido"})

	sort.Slice(events, func(i, j int) bool {
		return events[i].Minute < events[j].Minute
	})

	return events
}
func getRandomDefender(lineup []Lineup) *Lineup {
	defenders := filterPlayersByPosition(lineup, "Defender")
	return getRandomPlayer(defenders)
}

func getRandomMidfielder(lineup []Lineup) *Lineup {
	midfielders := filterPlayersByPosition(lineup, "Midfielder")
	return getRandomPlayer(midfielders)
}

func getRandomForward(lineup []Lineup) *Lineup {
	forwards := filterPlayersByPosition(lineup, "Forward")
	return getRandomPlayer(forwards)
}

func getRandomPlayer(players []Lineup) *Lineup {
	if len(players) == 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	return &players[rand.Intn(len(players))]
}
func getRandomPlayerExcludingGoalkeeper(lineup []Lineup) *Lineup {
	players := filterOutGoalkeepers(lineup)
	return getRandomPlayer(players)
}

func filterOutGoalkeepers(lineup []Lineup) []Lineup {
	var filtered []Lineup
	for _, player := range lineup {
		if player.Position != "Goalkeeper" {
			filtered = append(filtered, player)
		}
	}
	return filtered
}
