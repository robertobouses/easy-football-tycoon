package staff

import (
	//"database/sql"
	"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetStaffRandomByAnalytics(scouting int) (app.Staff, error) {

	row := r.getStaffRandomByAnalytics.QueryRow(scouting)
	log.Println("el valor del scouting es", scouting)
	var staff app.Staff
	if err := row.Scan(
		&staff.StaffId,
		&staff.StaffName,
		&staff.Job,
		&staff.Age,
		&staff.Fee,
		&staff.Salary,
		&staff.Training,
		&staff.Finances,
		&staff.Scouting,
		&staff.Physiotherapy,
		&staff.Rarity,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No se encontraron datos GetStaffRandomByAnalytics")
			return app.Staff{}, nil
		}
		log.Printf("Error al escanear la fila: %v", err)
		return app.Staff{}, err
	}

	return staff, nil
}
