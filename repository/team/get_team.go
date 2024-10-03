package team

import (
	//"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetTeam() ([]app.Player, error) {

	rows, err := r.getTeam.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []app.Player
	for rows.Next() {
		var player app.Player
		if err := rows.Scan(
			&player.PlayerId,
			&player.FirstName,
			&player.LastName,
			&player.Nationality,
			&player.Position,
			&player.Age,
			&player.Fee,
			&player.Salary,
			&player.Technique,
			&player.Mental,
			&player.Physique,
			&player.InjuryDays,
			&player.Lined,
			&player.Familiarity,
			&player.Fitness,
			&player.Happiness,
		); err != nil {
			log.Printf("Error al escanear las filas GetTeam: %v", err)
			return nil, err
		}
		log.Printf("Jugador recuperado de la base de datos: %+v", player)

		players = append(players, player)
	}
	log.Printf("Jugadores obtenido con éxito: %+v", players)
	return players, nil
}
