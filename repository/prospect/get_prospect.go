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

	var prospects []app.Prospect
	for rows.Next() {
		var prospect app.Prospect
		if err := rows.Scan(
			&prospect.ProspectId,
			&prospect.ProspectName,
			&prospect.Position,
			&prospect.Age,
			&prospect.Fee,
			&prospect.Salary,
			&prospect.Technique,
			&prospect.Mental,
			&prospect.Physique,
			&prospect.InjuryDays,
			&prospect.Job,
			&prospect.Rarity,
		); err != nil {
			log.Printf("Error al escanear las filas: %v", err)
			return nil, err
		}
		prospects = append(prospects, prospect)
	}

	return prospects, nil
}
