package app

import (
	"fmt"
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

	return numberOfMatchEvents, nil
}

func DistributeMatchEvents(lineup, rival []Lineup, numberOfMatchEvents int, homeOrAwayString string) (int, int, error) {
	lineupTotalQuality, err := CalculateQuality(lineup)
	if err != nil {
		return 0, 0, err
	}
	rivalTotalQuality, err := CalculateQuality(lineup)
	if err != nil {
		return 0, 0, err
	}
	allQuality := lineupTotalQuality + rivalTotalQuality
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

	rivalTypeLineup = append(rivalTypeLineup, Lineup{
		PlayerId:  uuid.New(),
		LastName:  "goalkeeper",
		Position:  "goalkeeper",
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
			LastName:  "defender" + string(i+1),
			Position:  "defender",
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
			LastName:  "midfielder" + string(i+1),
			Position:  "midfielder",
			Technique: clamp(rival.Technique+rand.Intn(21)-20, 0, 100),
			Mental:    clamp(rival.Mental+rand.Intn(21)-20, 0, 100),
			Physique:  clamp(rival.Physique+rand.Intn(21)-20, 0, 100),
		})
	}

	numForwards := 11 - playersAssigned
	for i := 0; i < numForwards; i++ {
		rivalTypeLineup = append(rivalTypeLineup, Lineup{
			PlayerId:  uuid.New(),
			LastName:  "forward" + string(i+1),
			Position:  "forward",
			Technique: clamp(rival.Technique+rand.Intn(21)-20, 0, 100),
			Mental:    clamp(rival.Mental+rand.Intn(3), 0, 100),
			Physique:  clamp(rival.Physique+rand.Intn(3), 0, 100),
		})
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
	Event           string `json:"event"`
	Minute          int    `json:"minute"`
	AttackOrDefense string `json:"attackordefense"`
}

func (a AppService) GenerateEvents(lineup, rivalLineup []Lineup, numberOfLineupEvents, numberOfRivalEvents int) ([]EventResult, []EventResult, int, int, int, int) {

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
	}
	var lineupResults []EventResult
	var rivalResults []EventResult
	var lineupScoreChances, rivalScoreChances, lineupGoals, rivalGoals int

	for i := 0; i < numberOfLineupEvents; i++ {
		event := lineupEvents[rand.Intn(len(lineupEvents))]
		result, _, _, _, _, err := event.Execute()
		minute := (i + 1) * 9
		if err != nil {
			lineupResults = append(lineupResults, EventResult{Event: fmt.Sprintf("Error en el evento: %v", err), Minute: minute, AttackOrDefense: "lineup"})
		} else {
			lineupResults = append(lineupResults, EventResult{Event: result, Minute: minute, AttackOrDefense: "lineup"})
		}
	}

	for i := 0; i < numberOfRivalEvents; i++ {
		event := rivalEvents[rand.Intn(len(rivalEvents))]
		result, _, _, _, _, err := event.Execute()
		minute := (i + 1) * 9
		if err != nil {
			rivalResults = append(rivalResults, EventResult{Event: fmt.Sprintf("Error en el evento: %v", err), Minute: minute, AttackOrDefense: "rival"})
		} else {
			rivalResults = append(rivalResults, EventResult{Event: result, Minute: minute, AttackOrDefense: "rival"})
		}
	}

	return lineupResults, rivalResults, lineupScoreChances, rivalScoreChances, lineupGoals, rivalGoals
}
