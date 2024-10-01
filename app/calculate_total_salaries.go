package app

func calculateTotalTeamSalary(players []Team) int {
	totalSalary := 0
	for _, staff := range players {
		totalSalary += staff.Salary
	}
	return totalSalary
}

func calculateTotalStaffSalary(employees []Staff) int {
	totalSalary := 0
	for _, player := range employees {
		totalSalary += player.Salary
	}
	return totalSalary
}
