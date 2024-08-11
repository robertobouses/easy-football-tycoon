package calendary

import (
	//"database/sql"
	"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetAnalytics() (app.Analytics, error) {

	row, err := r.GetAnalytics.QueryRow()
	var analytics app.Analytics
	if err := row.Scan(
		&analytics.AnalyticsId,
		&analytics.Finances,
		&analytics.Scouting,
		&analytics.Physiotherapy,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No se encontraron datos")
			return nil, nil
		}
		log.Printf("Error al escanear la fila: %v", err)
		return nil, err
	}

	return &analytics, nil
}
