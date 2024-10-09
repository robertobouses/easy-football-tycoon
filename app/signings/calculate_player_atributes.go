package signings

import (
	"log"
	"math/rand"
)

const (
	emergingPlayer    = 18
	veryYoungPlayer   = 21
	youngPlayer       = 24
	primePlayer       = 27
	experiencedPlayer = 31
	oldPlayer         = 33
	veryOldPlayer     = 35
)

func CalculatePlayerAtributes() (string, int, int, int, int, int) {

	var technique, mental, physique, injuryDays int
	position := positions[rand.Intn(len(positions))]
	age := rand.Intn(21) + 16

	log.Println("edad", age)

	switch {
	case age < emergingPlayer:
		log.Println("jugador tipo emergingPlayer")
		technique = rand.Intn(73) + 1
		mental = rand.Intn(61) + 1
		physique = rand.Intn(82) + 3

	case age < veryYoungPlayer:
		log.Println("jugador tipo veryYoungPlayer")
		technique = rand.Intn(80) + 1
		mental = rand.Intn(64) + 1
		physique = rand.Intn(73) + 23

	case age < youngPlayer:
		log.Println("jugador tipo youngPlayer")
		technique = rand.Intn(90) + 1
		mental = rand.Intn(76) + 1
		physique = rand.Intn(63) + 38

	case age < primePlayer:
		log.Println("jugador tipo primePlayer")
		technique = rand.Intn(85) + 16
		mental = rand.Intn(80) + 21
		physique = rand.Intn(68) + 33

	case age < experiencedPlayer:
		log.Println("jugador tipo experiencedPlayer")
		technique = rand.Intn(80) + 21
		mental = rand.Intn(70) + 31
		physique = rand.Intn(70) + 27

	case age < oldPlayer:
		log.Println("jugador tipo oldPlayer")
		technique = rand.Intn(80) + 21
		mental = rand.Intn(60) + 41
		physique = rand.Intn(70) + 14

	case age < veryOldPlayer:
		log.Println("jugador tipo veryOldPlayer")
		technique = rand.Intn(90) + 11
		mental = rand.Intn(50) + 51
		physique = rand.Intn(69) + 1

	default:
		log.Println("jugador tipo default/mas viejo que veryOldPlayer")
		technique = rand.Intn(70) + 4
		mental = rand.Intn(65) + 24
		physique = rand.Intn(40) + 1
	}

	log.Printf("Valores: Technique=%d, Mental=%d, Physique=%d", technique, mental, physique)

	if physique < 12 {
		injuryDays = rand.Intn(32)
	}
	if physique < 27 {
		injuryDays = rand.Intn(17)
	}

	log.Println("position, age, technique, mental, physique, injuryDays", position, age, technique, mental, physique, injuryDays)

	return position, age, technique, mental, physique, injuryDays
}

var positions = []string{"goalkeeper", "defender", "defender", "defender", "midfielder", "midfielder", "midfielder", "forward", "forward"}
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
