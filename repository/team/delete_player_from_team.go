package team

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) DeletePlayerFromTeam(player app.Player) error {
	log.Printf("Intentando eliminar jugador con PlayerId: %v", player.PlayerId)
	_, err := r.deletePlayerFromTeam.Exec(player.PlayerId)
	if err != nil {
		log.Printf("Error en deletePlayerFromTeam repo: %v", err)
		return err
	}
	log.Printf("El jugador %v se fue de tu equipo", player.LastName)
	return nil
}
