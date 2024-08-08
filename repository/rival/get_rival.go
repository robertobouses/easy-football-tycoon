package rival

import (
	//"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetRival() ([]app.Rival, error) {

	rows, err := r.getRival.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []app.Rival
	for rows.Next() {
		var team app.Rival
		if err := rows.Scan(
			&team.RivalId,
			&team.RivalName,
			&team.Technique,
			&team.Mental,
			&team.Physique,
		); err != nil {
			log.Printf("Error al escanear las filas: %v", err)
			return nil, err
		}
		teams = append(teams, team)
	}

	return teams, nil
}
