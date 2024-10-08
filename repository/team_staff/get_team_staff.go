package team_staff

import (
	//"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

func (r *repository) GetTeamStaff() ([]staff.Staff, error) {

	rows, err := r.getTeamStaff.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var staffs []staff.Staff
	for rows.Next() {
		var staff staff.Staff
		if err := rows.Scan(
			&staff.StaffId,
			&staff.FirstName,
			&staff.LastName,
			&staff.Nationality,
			&staff.Job,
			&staff.Age,
			&staff.Fee,
			&staff.Salary,
			&staff.Training,
			&staff.Finances,
			&staff.Scouting,
			&staff.Physiotherapy,
			&staff.Knowledge,
			&staff.Intelligence,
			&staff.Rarity,
		); err != nil {
			log.Printf("Error al escanear las filas GetTeamStaff: %v", err)
			return nil, err
		}
		log.Printf("Empleado del staff del equipo recuperado de la base de datos: %+v", staff)

		staffs = append(staffs, staff)
	}
	log.Printf("Empleados del staff del equipo obtenido con éxito: %+v", staffs)
	return staffs, nil
}
