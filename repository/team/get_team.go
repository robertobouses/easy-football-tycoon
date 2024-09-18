package team

import (
	//"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetTeam() ([]app.Team, error) {

	rows, err := r.getTeam.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []app.Team
	for rows.Next() {
		var team app.Team
		if err := rows.Scan(
			&team.PlayerId,
			&team.PlayerName,
			&team.Position,
			&team.Age,
			&team.Fee,
			&team.Salary,
			&team.Technique,
			&team.Mental,
			&team.Physique,
			&team.InjuryDays,
			&team.Lined,
			&team.Familiarity,
			&team.Fitness,
			&team.Happiness,
		); err != nil {
			log.Printf("Error al escanear las filas GetTeam: %v", err)
			return nil, err
		}
		log.Printf("Jugador recuperado de la base de datos: %+v", team)

		teams = append(teams, team)
	}
	log.Printf("Jugadores obtenido con éxito: %+v", teams)
	return teams, nil
}
