package signings

import (
	//"database/sql"
	"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetSigningsRandomByAnalytics(scouting int) (app.Signings, error) {

	row := r.getSigningsRandomByAnalytics.QueryRow(scouting)
	log.Println("el valor del scouting es", scouting)
	var signings app.Signings
	if err := row.Scan(
		&signings.SigningsId,
		&signings.FirstName,
		&signings.LastName,
		&signings.Nationality,
		&signings.Position,
		&signings.Age,
		&signings.Fee,
		&signings.Salary,
		&signings.Technique,
		&signings.Mental,
		&signings.Physique,
		&signings.InjuryDays,
		&signings.Rarity,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No se encontraron datos GetSigningsRandomByAnalytics")
			return app.Signings{}, nil
		}
		log.Printf("Error al escanear la fila: %v", err)
		return app.Signings{}, err
	}

	return signings, nil
}
