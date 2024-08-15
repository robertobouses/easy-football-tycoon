package prospect

import (
	//"database/sql"
	"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetProspectRandomByAnalytics(scouting int) (app.Prospect, error) {

	row := r.getProspectRandomByAnalytics.QueryRow()
	var prospect app.Prospect
	if err := row.Scan(
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
		if err == sql.ErrNoRows {
			log.Printf("No se encontraron datos")
			return app.Prospect{}, nil
		}
		log.Printf("Error al escanear la fila: %v", err)
		return app.Prospect{}, err
	}

	return prospect, nil
}
