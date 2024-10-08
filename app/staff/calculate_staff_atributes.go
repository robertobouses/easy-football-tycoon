package staff

import "math/rand"

func CalculateStaffAtributes() (string, int, int, int, int, int, int, int) {

	job := jobs[rand.Intn(len(jobs))]
	age := rand.Intn(30) + 33
	var training, finances, scouting, physiotherapy int
	switch job {
	case "trainer":
		training = rand.Intn(100) + 1
		finances = rand.Intn(50) + 1
		scouting = rand.Intn(95) + 1
		physiotherapy = rand.Intn(95) + 1

	case "financial":
		training = rand.Intn(20) + 1
		finances = rand.Intn(100) + 1
		scouting = rand.Intn(80) + 1
		physiotherapy = rand.Intn(40) + 1
	case "scout":
		training = rand.Intn(30) + 1
		finances = rand.Intn(70) + 1
		scouting = rand.Intn(100) + 1
		physiotherapy = rand.Intn(20) + 1
	case "physiotherapist":
		training = rand.Intn(80) + 1
		finances = rand.Intn(20) + 1
		scouting = rand.Intn(80) + 1
		physiotherapy = rand.Intn(100) + 1

	}
	knowledge := rand.Intn(100) + 1
	intelligence := rand.Intn(100) + 1

	return job, age, training, finances, scouting, physiotherapy, knowledge, intelligence
}

var jobs = []string{"trainer", "financial", "scout", "physiotherapist"}
var nationalities = []string{
	"gb", "gb", "gb", "gb", "gb", "gb", "gb", "gb",
	"es", "es", "es", "es", "es", "es", "es",
	"fr", "fr", "fr", "fr", "fr", "fr",
	"de", "de", "de", "de", "de", "de",
	"it", "it", "it", "it", "it", "it",
	"nl", "nl", "nl", "nl", "nl", "nl",
	"pt", "pt", "pt", "pt",

	"ar", "br", "mx", "jp", "kr", "cn", "ru", "se", "no", "dk",
	"au", "ca", "pl", "be", "ch", "at", "fi", "cz", "gr", "ro",
	"ie", "hu", "sk", "hr", "tr", "za", "co", "uy", "cl", "ec",
}
