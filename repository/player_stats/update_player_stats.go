package player_stats

import (
	"log"

	"github.com/google/uuid"
)

func (r *repository) UpdatePlayerStats(playerId uuid.UUID, appearances int, blocks int, saves int, aerialDuel int, keyPass int, assists int, chances int, goals int, mvp int, rating float64) error {
	log.Printf("Updating player data: %v, %v, %v, %v, %v, %v, %v",
		playerId, appearances, blocks, saves, aerialDuel, chances, assists, goals, mvp, rating)

	_, err := r.updatePlayerStats.Exec(

		appearances,
		blocks,
		saves,
		aerialDuel,
		keyPass,
		assists,
		chances,
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
