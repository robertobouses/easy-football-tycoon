package staff

import (
	"math/rand"
	"time"
)

const (
	veryHighStaffQuality = 5690
	highStaffQuality     = 4820
	mediumStaffQuality   = 3150
	lowStaffQuality      = 730
	veryOldAge           = 55
	oldAge               = 45
	trainerBonus         = 2
	financialBonus       = 3
)

func CalculateStaffFeeAndSalary(training, finances, scouting, physiotherapy, knowledge, intelligence, age int, nat, job string) (int, int, int) {
	rand.Seed(time.Now().UnixNano())

	totalStaffQuality := training + finances + scouting + physiotherapy + knowledge + intelligence

	var totalSpecialistQuality int

	switch job {
	case "trainer":
		totalSpecialistQuality = 10*training + scouting + physiotherapy + 5*knowledge + 3*intelligence

	case "financial":
		totalSpecialistQuality = 12*finances + scouting + 5*knowledge + 2*intelligence

	case "scout":
		totalSpecialistQuality = finances + 12*scouting + 2*knowledge + 5*intelligence

	case "physiotherapist":
		totalSpecialistQuality = training + scouting + 10*physiotherapy + 5*knowledge + 3*intelligence

	default:
		totalSpecialistQuality = 0
	}

	staffQuality := totalStaffQuality + 3*totalSpecialistQuality
	salaryBase := staffQuality * 2

	var salary, fee, rarity int

	switch {
	case staffQuality > veryHighStaffQuality:
		salary = int(float64(10) * float64(salaryBase) * (rand.Float64()*3.1 + 1))
		fee = 1_000_000 * (rand.Intn(4) + 1)
		rarity = rand.Intn(29) + 72

	case staffQuality > highStaffQuality:
		salary = int(float64(5) * float64(salaryBase) * (rand.Float64()*3.4 + 1))
		fee = 500_000 * (rand.Intn(3) + 1)
		rarity = rand.Intn(41) + 50

	case staffQuality > mediumStaffQuality:
		salary = int(float64(3) * float64(salaryBase) * (rand.Float64()*5.3 + 1))
		fee = 5_000 * (rand.Intn(34) + 6)
		rarity = rand.Intn(31) + 40

	case staffQuality > lowStaffQuality:
		salary = int(float64(2) * float64(salaryBase) * (rand.Float64()*3.7 + 1))
		fee = 5_000 * (rand.Intn(14) + 1)
		rarity = rand.Intn(31) + 20

	default:
		fee = 5_000 * (rand.Intn(16) + 1)
		salary = int(float64(2) * float64(salaryBase) * (rand.Float64()*2.9 + 1))
		rarity = rand.Intn(19) + 1
	}

	switch {
	case age > veryOldAge:
		fee -= 150_000 * (rand.Intn(4) + 1)
		salary += 4_000 * (7 * age / veryOldAge)
	case age > oldAge:
		salary += 2_000 * (4 * age / oldAge)
	}

	switch nat {
	case "gb", "de":
		salary += 1_000 * (rand.Intn(9) + 2)

	}

	if job == "financial" {
		salary *= financialBonus
	}

	if job == "trainer" {
		salary *= trainerBonus
	}

	if fee < 0 {
		fee = 0
	}
	if salary < 21_000 {
		salary = 21_000
	}
	if rarity < 0 {
		rarity = 1
	}

	return fee, salary, rarity
}
