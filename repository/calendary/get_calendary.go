package calendary

import (
	//"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetCalendary() ([]app.Calendary, error) {

	rows, err := r.getCalendary.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var completeCalendary []app.Calendary
	for rows.Next() {
		var daytype app.Calendary
		if err := rows.Scan(
			&daytype.DayType,
		); err != nil {
			log.Printf("Error al escanear las filas: %v", err)
			return nil, err
		}
		completeCalendary = append(completeCalendary, daytype)
	}

	return completeCalendary, nil
}
