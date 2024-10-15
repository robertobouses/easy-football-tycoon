package player_stats

import (
	"log"

	"github.com/google/uuid"
)

func (r *repository) UpdatePlayerStats(playerId uuid.UUID, appearances int, chances int, assists int, goals int, mvp int, rating float64) error {
	log.Printf("Updating player data: %v, %v, %v, %v, %v, %v, %v",
		playerId, appearances, chances, assists, goals, mvp, rating)

	_, err := r.updatePlayerStats.Exec(

		appearances,
		chances,
		assists,
		goals,
		mvp,
		rating,
		playerId,
	)
	if err != nil {
		log.Printf("Error en updatePlayerLinedStatus repo: %v", err)
		return err
	}
	log.Println("Datos del jugador actualizados correctamente")
	return nil
}
