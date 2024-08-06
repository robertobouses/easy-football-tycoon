package team

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) PostProspect(req app.Prospect) error {
	_, err := r.postProspect.Exec(
		req.ProspectId,
		req.ProspectName,
		req.Position,
		req.Age,
		req.Fee,
		req.Salary,
		req.Technique,
		req.Mental,
		req.Physique,
		req.InjuryDays,
		req.Lined,
		req.Job,
		req.Rarity,
	)

	if err != nil {
		log.Print("Error en PostProspect repo", err)
		return err
	}
	log.Println("Después de ejecutar la consulta preparada")
	return nil
}
