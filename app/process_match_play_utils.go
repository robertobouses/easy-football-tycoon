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

func (a *AppService) GetRandomDefender(lineup []Lineup) *Lineup {
	defenders := filterPlayersByPosition(lineup, "Defender")
	return GetRandomPlayer(defenders)
}

func (a *AppService) GetRandomMidfielder(lineup []Lineup) *Lineup {
	midfielders := filterPlayersByPosition(lineup, "Midfielder")
	return GetRandomPlayer(midfielders)
}

func (a *AppService) GetRandomForward(lineup []Lineup) *Lineup {
	forwards := filterPlayersByPosition(lineup, "Forward")
	return GetRandomPlayer(forwards)
}

func (a *AppService) GetGoalkeeper(lineup []Lineup) *Lineup {
	goalkeepers := filterPlayersByPosition(lineup, "Goalkeeper")
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
		if player.Position != "Goalkeeper" {
			filtered = append(filtered, player)
		}
	}
	return filtered
}

type Event struct {
	Name    string
	Execute func() (string, error)
}

func (a AppService) GenerateEvents(lineup, rivalLineup []Lineup, numberOfLineupEvents, numberOfRivalEvents int) {

	lineupEvents := []Event{
		{
			"Pase clave",
			func() (string, error) {
				return a.KeyPass(lineup, rivalLineup)
			},
		},
		{
			"Remate a puerta",
			func() (string, error) {
				return a.Shot(lineup, rivalLineup, a.GetRandomForward(lineup))
			},
		},
		{
			"Penalty",
			func() (string, error) {
				return a.PenaltyKick(lineup, rivalLineup)
			},
		},
		{
			"Tiro lejano",
			func() (string, error) {
				return a.LongShot(lineup, rivalLineup)
			},
		},
	}

	rivalEvents := []Event{
		{
			"Pase clave",
			func() (string, error) {
				return a.KeyPass(rivalLineup, lineup)
			},
		},
		{
			"Remate a puerta",
			func() (string, error) {
				return a.Shot(rivalLineup, lineup, a.GetRandomForward(rivalLineup))
			},
		},
		{
			"Penalty",
			func() (string, error) {
				return a.PenaltyKick(rivalLineup, lineup)
			},
		},
		{
			"Tiro lejano",
			func() (string, error) {
				return a.LongShot(rivalLineup, lineup)
			},
		},
	}

	for i := 0; i < numberOfLineupEvents; i++ {
		evento := lineupEvents[rand.Intn(len(lineupEvents))]
		fmt.Printf("Minuto %d: ", (i+1)*9)

		result, err := evento.Execute()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(result)
		}
	}

	fmt.Println("\n--- Eventos del rival ---")
	for i := 0; i < numberOfRivalEvents; i++ {
		evento := rivalEvents[rand.Intn(len(rivalEvents))]
		fmt.Printf("Minuto %d: ", (i+1)*9)

		result, err := evento.Execute()
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(result)
		}
	}
}
