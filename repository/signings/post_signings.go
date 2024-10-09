package signings

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app/signings"
)

func (r *repository) PostSignings(req signings.Signings) error {
	_, err := r.postSignings.Exec(
		req.FirstName,
		req.LastName,
		req.Nationality,
		req.Position,
		req.Age,
		req.Fee,
		req.Salary,
		req.Technique,
		req.Mental,
		req.Physique,
		req.InjuryDays,
		req.Rarity,
		req.Fitness,
	)

	if err != nil {
		log.Print("Error en PostSignings repo", err)
		return err
	}
	log.Println("Después de ejecutar la consulta preparada")
	return nil
}
