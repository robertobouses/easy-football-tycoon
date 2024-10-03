package lineup

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) PostLineup(player app.Player) error {
	_, err := r.postLineup.Exec(
		player.PlayerId,
		player.LastName,
		player.Position,
		player.Technique,
		player.Mental,
		player.Physique,
	)

	if err != nil {
		log.Print("Error en PostLineup repo", err)
		return err
	}
	log.Println("Después de ejecutar la consulta preparada")
	return nil
}
