package team

import (
	//"database/sql"
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetPlayerByPlayerId(playerId uuid.UUID) (*app.Team, error) {

	row := r.getPlayerByPlayerId.QueryRow(playerId)
	var player app.Team
	if err := row.Scan(
		&player.PlayerId,
		&player.PlayerName,
		&player.Position,
		&player.Technique,
		&player.Mental,
		&player.Physique,
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
