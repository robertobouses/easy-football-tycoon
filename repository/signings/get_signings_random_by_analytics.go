package signings

import (
	//"database/sql"
	"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app/signings"
)

func (r *repository) GetSigningsRandomByAnalytics(scouting int) (signings.Signings, error) {

	row := r.getSigningsRandomByAnalytics.QueryRow(scouting)
	log.Println("el valor del scouting es", scouting)
	var signing signings.Signings
	if err := row.Scan(
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
		if err == sql.ErrNoRows {
			log.Printf("No se encontraron datos GetSigningsRandomByAnalytics")
			return signings.Signings{}, nil
		}
		log.Printf("Error al escanear la fila: %v", err)
		return signings.Signings{}, err
	}

	return signing, nil
}
