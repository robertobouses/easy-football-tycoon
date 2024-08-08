package rival

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) PostRival(req app.Rival) error {
	_, err := r.postRival.Exec(
		req.RivalName,
		req.Technique,
		req.Mental,
		req.Physique,
	)

	if err != nil {
		log.Print("Error en PostRival repo", err)
		return err
	}
	log.Println("Después de ejecutar la consulta preparada")
	return nil
}
