package team

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) UpdatePlayerLinedStatus(req app.Lineup) error {
	log.Println("Antes de ejecutar la consulta preparada")
	_, err := r.updatePlayerLinedStatus.Exec(
		req.PlayerId,
	)

	if err != nil {
		log.Print("Error en updatePlayerLinedStatus repo", err)
		return err
	}
	log.Println("Después de ejecutar la consulta preparada")
	return nil
}
