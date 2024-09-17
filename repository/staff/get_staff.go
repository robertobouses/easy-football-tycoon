package staff

import (
	//"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetStaff() ([]app.Staff, error) {

	rows, err := r.getStaff.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var staffs []app.Staff
	for rows.Next() {
		var staff app.Staff
		if err := rows.Scan(
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
			log.Printf("Error al escanear las filas GetStaff: %v", err)
			return nil, err
		}
		log.Printf("Empleado recuperado de la base de datos: %+v", staff)

		staffs = append(staffs, staff)
	}
	log.Printf("Empleado obtenido con éxito: %+v", staffs)
	return staffs, nil
}
