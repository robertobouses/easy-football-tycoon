package app

import "github.com/robertobouses/easy-football-tycoon/app/staff"

func calculateTotalTeamSalary(players []Player) int {
	totalSalary := 0
	for _, staff := range players {
		totalSalary += staff.Salary
	}
	return totalSalary
}

func calculateTotalStaffSalary(employees []staff.Staff) int {
	totalSalary := 0
	for _, player := range employees {
		totalSalary += player.Salary
	}
	return totalSalary
}
