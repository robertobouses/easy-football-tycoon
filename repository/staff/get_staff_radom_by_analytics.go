package staff

import (
	//"database/sql"
	"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

func (r *repository) GetStaffRandomByAnalytics(scouting int) (staff.Staff, error) {

	row := r.getStaffRandomByAnalytics.QueryRow(scouting)
	log.Println("el valor del scouting es", scouting)
	var oneStaff staff.Staff
	if err := row.Scan(
		&oneStaff.StaffId,
		&oneStaff.FirstName,
		&oneStaff.LastName,
		&oneStaff.Nationality,
		&oneStaff.Job,
		&oneStaff.Age,
		&oneStaff.Fee,
		&oneStaff.Salary,
		&oneStaff.Training,
		&oneStaff.Finances,
		&oneStaff.Scouting,
		&oneStaff.Physiotherapy,
		&oneStaff.Knowledge,
		&oneStaff.Intelligence,
		&oneStaff.Rarity,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No se encontraron datos GetStaffRandomByAnalytics")
			return staff.Staff{}, nil
		}
		log.Printf("Error al escanear la fila: %v", err)
		return staff.Staff{}, err
	}

	return oneStaff, nil
}
