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

	var rivals []app.Rival
	for rows.Next() {
		var rival app.Rival
		if err := rows.Scan(
			&rival.RivalId,
			&rival.RivalName,
			&rival.Technique,
			&rival.Mental,
			&rival.Physique,
		); err != nil {
			log.Printf("Error al escanear las filas: %v", err)
			return nil, err
		}
		rivals = append(rivals, rival)
	}

	return rivals, nil
}
