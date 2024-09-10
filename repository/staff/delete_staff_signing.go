package staff

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) DeleteStaffSigning(staff app.Staff) error {
	log.Printf("Intentando eliminar Empleado con staffId: %v", staff.StaffId)
	_, err := r.deleteStaffSigning.Exec(staff.StaffId)
	if err != nil {
		log.Printf("Error en deletestaffFromTeam repo: %v", err)
		return err
	}
	log.Printf("El Empleado %v se fue de tu equipo", staff.StaffName)
	return nil
}
