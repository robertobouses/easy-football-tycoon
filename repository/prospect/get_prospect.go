package prospect

import (
	//"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetProspect() ([]app.Prospect, error) {

	rows, err := r.getProspect.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []app.Prospect
	for rows.Next() {
		var team app.Prospect
		if err := rows.Scan(
			&team.Fee,
			&team.ProspectId,
			&team.ProspectName,
			&team.Position,
			&team.Age,
			&team.Salary,
			&team.Technique,
			&team.Mental,
			&team.Physique,
			&team.InjuryDays,
			&team.Lined,
			&team.Job,
			&team.Rarity,
		); err != nil {
			log.Printf("Error al escanear las filas: %v", err)
			return nil, err
		}
		teams = append(teams, team)
	}

	return teams, nil
}
