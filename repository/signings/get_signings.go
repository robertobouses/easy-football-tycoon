package signings

import (
	//"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app/signings"
)

func (r *repository) GetSignings() ([]signings.Signings, error) {

	rows, err := r.getSignings.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var allSignings []signings.Signings
	for rows.Next() {
		var signing signings.Signings
		if err := rows.Scan(
			&signing.SigningsId,
			&signing.FirstName,
			&signing.LastName,
			&signing.Nationality,
			&signing.Position,
			&signing.Age,
			&signing.Fee,
			&signing.Salary,
			&signing.Technique,
			&signing.Mental,
			&signing.Physique,
			&signing.InjuryDays,
			&signing.Rarity,
		); err != nil {
			log.Printf("Error al escanear las filas GetSignings: %v", err)
			return nil, err
		}
		allSignings = append(allSignings, signing)
	}

	return allSignings, nil
}
