package staff

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

func (r *repository) PostStaff(req staff.Staff) error {
	_, err := r.postStaff.Exec(
		req.FirstName,
		req.LastName,
		req.Nationality,
		req.Job,
		req.Age,
		req.Fee,
		req.Salary,
		req.Training,
		req.Finances,
		req.Scouting,
		req.Physiotherapy,
		req.Knowledge,
		req.Intelligence,
		req.Rarity,
	)

	if err != nil {
		log.Print("Error en PostStaff repo", err)
		return err
	}
	log.Println("Después de ejecutar la consulta preparada")
	return nil
}
