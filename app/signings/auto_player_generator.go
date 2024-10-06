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

var positions = []string{"goalkeeper", "defender", "midfielder", "forward"}
var nationalities = []string{
	"gb", "gb", "gb", "gb", "gb",
	"es", "es", "es", "es", "es",
	"fr", "fr", "fr", "fr", "fr",
	"de", "de", "de", "de", "de",
	"it", "it", "it", "it", "it",
	"nl", "nl", "nl", "nl", "nl",
	"pt", "pt", "pt", "pt", "pt",

	"ar", "br", "mx", "jp", "kr", "cn", "ru", "se", "no", "dk",
	"au", "ca", "pl", "be", "ch", "at", "fi", "cz", "gr", "ro",
	"ie", "hu", "sk", "hr", "tr", "za", "co", "uy", "cl", "ec",
}

func (a *SigningsService) RunAutoPlayerGenerator(numberOfPlayers int) ([]Signings, error) {
	rand.Seed(time.Now().UnixNano())
	var players []Signings
	log.Printf("Número de jugadores a generar: %d", numberOfPlayers)

	for i := 0; i < numberOfPlayers; i++ {
		nat := nationalities[rand.Intn(len(nationalities))]

		firstName, lastName, _, err := getRandomNameByNationality(nat)
		if err != nil {
			return nil, fmt.Errorf("error generating player name: %v", err)
		}

		// Genera los atributos antes de calcular el salario y el fee
		age := rand.Intn(18) + 18
		technique := rand.Intn(100) + 1
		mental := rand.Intn(100) + 1
		physique := rand.Intn(100) + 1

		// Ahora sí puedes calcular el fee y salary
		fee, salary := CalculatePlayerFeeAndSalary(technique, mental, physique, age)

		log.Println("valor de age", age)
		log.Println("valor de technique", technique)
		log.Println("valor de mental", mental)
		log.Println("valor de physique", physique)

		player := Signings{
			FirstName:   firstName,
			LastName:    lastName,
			Nationality: nat,
			Position:    positions[rand.Intn(len(positions))],
			Age:         age,
			Fee:         fee,    // Usa el valor calculado
			Salary:      salary, // Usa el valor calculado
			Technique:   technique,
			Mental:      mental,
			Physique:    physique,
			InjuryDays:  rand.Intn(30),
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
