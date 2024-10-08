package team_staff

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

func (r *repository) DeleteTeamStaff(staff staff.Staff) error {
	log.Printf("Intentando eliminar jugador con StaffId: %v", staff.StaffId)
	_, err := r.deleteTeamStaff.Exec(staff.StaffId)
	if err != nil {
		log.Printf("Error en deleteStaffFromTeam repo: %v", err)
		return err
	}
	log.Printf("El empleado %v %v se fue de tu equipo", staff.FirstName, staff.LastName)
	return nil
}
