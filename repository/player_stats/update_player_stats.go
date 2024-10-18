package player_stats

import (
	"log"

	"github.com/google/uuid"
)

func (r *repository) UpdatePlayerStats(playerId uuid.UUID, appearances int, blocks int, saves int, aerialDuel int, keyPass int, assists int, chances int, goals int, mvp int, rating float64) error {
	log.Printf("Updating player data: %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v", playerId, appearances, blocks, saves, aerialDuel, keyPass, assists, chances, goals, mvp, rating)

	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM eft.lineup WHERE player_id = $1)", playerId).Scan(&exists)
	if err != nil {
		log.Printf("Error checking player existence: %v", err)
		return err
	}

	if !exists {
		log.Printf("Player with ID %v does not exist, skipping update.", playerId)
		return nil
	}
	log.Printf("Executing query with parameters: %v", []interface{}{
		playerId, appearances, blocks, saves, aerialDuel, keyPass, assists, chances, goals, mvp, rating,
	})
	_, err = r.updatePlayerStats.Exec(
		playerId,
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
	)
	if err != nil {
		log.Printf("Error en updatePlayerStats repo: %v, Query: %s, Params: %v", err, "INSERT INTO eft.player_stats (player_id, appearances, blocks, saves, aerialduel, keypass, assists, chances, goals, mvp, rating) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) ON CONFLICT (player_id) DO UPDATE SET ...", []interface{}{playerId, appearances, blocks, saves, aerialDuel, keyPass, assists, chances, goals, mvp, rating})
		return err
	}
	log.Println("Datos del jugador actualizados correctamente")
	return nil
}
