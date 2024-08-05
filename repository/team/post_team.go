package team

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) PostTeam(req app.Team) error {
	_, err := r.postTeam.Exec(
		req.PlayerName,
		req.Position,
		req.Age,
		req.Fee,
		req.Salary,
		req.Technique,
		req.Mental,
		req.Physique,
		req.InjuryDays,
		req.Lined,
	)

	if err != nil {
		log.Print("Error en PostTeam repo", err)
		return err
	}
	log.Println("Después de ejecutar la consulta preparada")
	return nil
}
