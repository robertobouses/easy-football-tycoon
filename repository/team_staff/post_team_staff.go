package team_staff

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

func (r *repository) PostTeamStaff(req staff.Staff) error {
	_, err := r.postTeamStaff.Exec(
		req.StaffName,
		req.Job,
		req.Age,
		req.Fee,
		req.Salary,
		req.Training,
		req.Finances,
		req.Scouting,
		req.Physiotherapy,
		req.Rarity,
	)

	if err != nil {
		log.Print("Error en PostTeamStaff repo", err)
		return err
	}
	log.Println("Después de ejecutar la consulta preparada PostTeamStaff")
	return nil
}
