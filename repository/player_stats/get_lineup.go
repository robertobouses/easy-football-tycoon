package player_stats

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetPlayerStats() ([]app.PlayerStats, error) {

	rows, err := r.getPlayerStats.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var playerStatss []app.PlayerStats
	for rows.Next() {
		var playerStats app.PlayerStats
		if err := rows.Scan(
			&playerStats.PlayerId,
			&playerStats.Appearances,
			&playerStats.Chances,
			&playerStats.Assists,
			&playerStats.Goals,
			&playerStats.MVP,
			&playerStats.Rating,
		); err != nil {
			log.Printf("Error al escanear las filas GetPlayerStats: %v", err)
			return nil, err
		}
		playerStatss = append(playerStatss, playerStats)

		if len(playerStatss) >= 11 {
			break
		}
	}

	return playerStatss, nil
}
