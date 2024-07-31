package team

import (
	"database/sql"
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
		); err != nil {
			log.Printf("Error al escanear las filas: %v", err)
			return nil, err
		}
		teams = append(teams, team)
	}

	return teams, nil
}

type repository struct {
	db      *sql.DB
	getTeam *sql.Stmt
}
