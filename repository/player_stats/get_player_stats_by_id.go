package player_stats

import (
	//"database/sql"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetPlayerStatsById(playerId uuid.UUID) (*app.PlayerStats, error) {

	row := r.getPlayerStatsById.QueryRow(playerId)
	var player app.PlayerStats
	if err := row.Scan(
		&player.PlayerId,
		&player.Appearances,
		&player.Blocks,
		&player.Saves,
		&player.AerialDuel,
		&player.KeyPass,
		&player.Assists,
		&player.Chances,
		&player.Goals,
		&player.MVP,
		&player.Rating,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No se encontró el jugador con playerid: %v", playerId)
			return nil, nil
		}
		log.Printf("Error al escanear la fila: %v", err)
		return nil, err
	}

	return &player, nil
}
