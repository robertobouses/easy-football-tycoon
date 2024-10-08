package staff

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/robertobouses/easy-football-tycoon/app/signings"
)

func (a *StaffService) RunAutoStaffGenerator(numberOfStaffs int) ([]Staff, error) {
	rand.Seed(time.Now().UnixNano())
	var staffs []Staff
	log.Printf("Número de empleados a generar: %d", numberOfStaffs)

	for i := 0; i < numberOfStaffs; i++ {
		nat := nationalities[rand.Intn(len(nationalities))]

		firstName, lastName, _, err := signings.GetRandomNameByNationality(nat)
		if err != nil {
			return nil, fmt.Errorf("error generating staff name: %v", err)
		}

		job, age, training, finances, scouting, physiotherapy, knowledge, intelligence := CalculateStaffAtributes()

		fee, salary, rarity := CalculateStaffFeeAndSalary(training, finances, scouting, physiotherapy, knowledge, intelligence, age, nat, job)

		staff := Staff{
			FirstName:     firstName,
			LastName:      lastName,
			Nationality:   nat,
			Job:           job,
			Age:           age,
			Fee:           fee,
			Salary:        salary,
			Training:      training,
			Finances:      finances,
			Scouting:      scouting,
			Physiotherapy: physiotherapy,
			Knowledge:     knowledge,
			Intelligence:  intelligence,
			Rarity:        rarity,
		}
		staffs = append(staffs, staff)

		log.Println("jugador generado automáticamente", staff)

		err = a.staffRepo.PostStaff(staff)
		if err != nil {
			log.Printf("Error insertando jugador %v %v: %v", staff.FirstName, staff.LastName, err)
		} else {
			log.Printf("Jugador %v %v insertado correctamente", staff.FirstName, staff.LastName)
		}
	}
	log.Println("todos los jugadores generados", staffs)
	return staffs, nil
}

// func getRandomNameByNationality(nationality string) (string, string, string, error) {
// 	baseURL := "https://randomuser.me/api/"
// 	params := url.Values{}
// 	params.Add("nat", nationality)
// 	params.Add("gender", "male")
// 	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

// 	resp, err := http.Get(apiURL)
// 	if err != nil {
// 		log.Printf("Error al obtener nombre de la API: %v", err)
// 		return "", "", "", err
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Printf("Error al leer el cuerpo de la respuesta de la API: %v", err)
// 		return "", "", "", err
// 	}

// 	log.Printf("Respuesta de la API: %s", string(body))

// 	var userName RandomUserName
// 	if err := json.Unmarshal(body, &userName); err != nil {
// 		log.Printf("Error al deserializar la respuesta JSON: %v", err)
// 		return "", "", "", err
// 	}

// 	firstName := userName.Results[0].Name.First
// 	lastName := userName.Results[0].Name.Last
// 	nationalityCode := userName.Results[0].Nat

// 	return firstName, lastName, nationalityCode, nil
// }

// type RandomUserName struct {
// 	Results []struct {
// 		Name struct {
// 			First string `json:"first"`
// 			Last  string `json:"last"`
// 		} `json:"name"`
// 		Nat string `json:"nat"`
// 	} `json:"results"`
// }
