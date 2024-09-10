package team_staff

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) DeleteTeamStaff(staff app.Staff) error {
	log.Printf("Intentando eliminar jugador con StaffId: %v", staff.StaffId)
	_, err := r.deleteTeamStaff.Exec(staff.StaffId)
	if err != nil {
		log.Printf("Error en deleteStaffFromTeam repo: %v", err)
		return err
	}
	log.Printf("El jugador %v se fue de tu equipo", staff.StaffName)
	return nil
}
