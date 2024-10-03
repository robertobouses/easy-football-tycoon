package autoplayergenerator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/robertobouses/easy-football-tycoon/app"
)

var positions = []string{"Goalkeeper", "Defender", "Midfielder", "Forward"}
var nationalities = []string{"us", "gb", "es", "fr", "de"} //TODO AMPLIAR NUMERO NAC. HACER MAS PROBABLE Y MAS POSIBLE CALIDAD CIERTAS NAC

func AutoPlayerGenerator(numberOfPlayers int) ([]app.Player, error) {
	rand.Seed(time.Now().UnixNano())
	var players []app.Player

	for i := 0; i < numberOfPlayers; i++ {
		nat := nationalities[rand.Intn(len(nationalities))]

		firstName, lastName, _, err := getRandomNameByNationality(nat)
		if err != nil {
			return nil, fmt.Errorf("error generating player name: %v", err)
		}

		player := app.Player{
			PlayerId:    uuid.New(),
			FirstName:   firstName,
			LastName:    lastName,
			Nationality: nat,
			Position:    positions[rand.Intn(len(positions))],
			Age:         rand.Intn(18) + 18,
			Fee:         rand.Intn(50_000_000) + 1_000_000, // TODO CALCULAR
			Salary:      rand.Intn(200_000) + 20_000,       // TODO CALCULAR
			Technique:   rand.Intn(100) + 1,
			Mental:      rand.Intn(100) + 1,
			Physique:    rand.Intn(100) + 1,
			InjuryDays:  rand.Intn(30),
			Lined:       false,
			Familiarity: rand.Intn(100) + 1,
			Fitness:     rand.Intn(100) + 1,
			Happiness:   rand.Intn(100) + 1,
		}
		players = append(players, player)
	}
	return players, nil
}

type RandomUserName struct {
	Results []struct {
		Name struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		Nat string `json:"nat"`
	} `json:"results"`
}

func getRandomNameByNationality(nationality string) (string, string, string, error) {
	baseURL := "https://randomuser.me/api/"
	params := url.Values{}
	params.Add("nat", nationality)
	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", err
	}

	var userName RandomUserName
	if err := json.Unmarshal(body, &userName); err != nil {
		return "", "", "", err
	}

	firstName := userName.Results[0].Name.First
	lastName := userName.Results[0].Name.Last
	nationalityCode := userName.Results[0].Nat

	return firstName, lastName, nationalityCode, nil
}
