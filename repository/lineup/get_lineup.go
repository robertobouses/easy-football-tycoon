package lineup

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetLineup() ([]app.Lineup, error) {

	rows, err := r.getLineup.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lineups []app.Lineup
	for rows.Next() {
		var lineup app.Lineup
		if err := rows.Scan(
			&lineup.PlayerName,
			&lineup.Position,
			&lineup.Technique,
			&lineup.Mental,
			&lineup.Physique,
			&lineup.TotalTechnique,
			&lineup.TotalMental,
			&lineup.TotalPhysique,
		); err != nil {
			log.Printf("Error al escanear las filas: %v", err)
			return nil, err
		}
		lineups = append(lineups, lineup)

		if len(lineups) >= 11 {
			break
		}
	}

	return lineups, nil
}
