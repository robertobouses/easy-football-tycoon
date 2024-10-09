package signings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func (a *SigningsService) RunAutoPlayerGenerator(numberOfPlayers int) ([]Signings, error) {
	rand.Seed(time.Now().UnixNano())
	var players []Signings
	log.Printf("Número de jugadores a generar: %d", numberOfPlayers)

	for i := 0; i < numberOfPlayers; i++ {
		nat := nationalities[rand.Intn(len(nationalities))]

		firstName, lastName, _, err := GetRandomNameByNationality(nat)
		if err != nil {
			return nil, fmt.Errorf("error generating player name: %v", err)
		}

		position, age, technique, mental, physique, injuryDays := CalculatePlayerAtributes()
		fee, salary, rarity := CalculatePlayerFeeAndSalary(technique, mental, physique, age, nat, position)

		log.Println("valor de age", age)
		log.Println("valor de technique", technique)
		log.Println("valor de mental", mental)
		log.Println("valor de physique", physique)

		player := Signings{
			FirstName:   firstName,
			LastName:    lastName,
			Nationality: nat,
			Position:    position,
			Age:         age,
			Fee:         fee,
			Salary:      salary,
			Technique:   technique,
			Mental:      mental,
			Physique:    physique,
			InjuryDays:  injuryDays,
			Rarity:      rarity,
			Fitness:     rand.Intn(100) + 1,
		}
		players = append(players, player)

		log.Println("jugador generado automáticamente", player)

		err = a.signingsRepo.PostSignings(player)
		if err != nil {
			log.Printf("Error insertando jugador %v %v: %v", player.FirstName, player.LastName, err)
		} else {
			log.Printf("Jugador %v %v insertado correctamente", player.FirstName, player.LastName)
		}
	}
	log.Println("todos los jugadores generados", players)
	return players, nil
}

func GetRandomNameByNationality(nationality string) (string, string, string, error) {
	baseURL := "https://randomuser.me/api/"
	params := url.Values{}
	params.Add("nat", nationality)
	params.Add("gender", "male")
	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(apiURL)
	if err != nil {
		log.Printf("Error al obtener nombre de la API: %v", err)
		return "", "", "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error al leer el cuerpo de la respuesta de la API: %v", err)
		return "", "", "", err
	}

	log.Printf("Respuesta de la API: %s", string(body))

	var userName RandomUserName
	if err := json.Unmarshal(body, &userName); err != nil {
		log.Printf("Error al deserializar la respuesta JSON: %v", err)
		return "", "", "", err
	}

	firstName := userName.Results[0].Name.First
	lastName := userName.Results[0].Name.Last
	nationalityCode := userName.Results[0].Nat

	return firstName, lastName, nationalityCode, nil
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
