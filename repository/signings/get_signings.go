package signings

import (
	//"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetSignings() ([]app.Signings, error) {

	rows, err := r.getSignings.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var signingss []app.Signings
	for rows.Next() {
		var signings app.Signings
		if err := rows.Scan(
			&signings.SigningsId,
			&signings.SigningsName,
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
			log.Printf("Error al escanear las filas GetSignings: %v", err)
			return nil, err
		}
		signingss = append(signingss, signings)
	}

	return signingss, nil
}
