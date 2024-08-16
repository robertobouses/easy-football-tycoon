package team

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) DeletePlayerFromTeam(player app.Team) error {
	_, err := r.deletePlayerFromTeam.Exec(player.PlayerId)
	if err != nil {
		log.Printf("Error en deletePlayerFromTeam repo: %v", err)
		return err
	}
	log.Println("El jugador %v se fue de tu equipo", player.PlayerName)
	return nil
}
