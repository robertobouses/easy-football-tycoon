package analytics

import (
	"database/sql"
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetAnalytics() (app.Analytics, error) {

	row := r.getAnalytics.QueryRow()
	var analytics app.Analytics
	if err := row.Scan(
		&analytics.AnalyticsId,
		&analytics.Training,
		&analytics.Finances,
		&analytics.Scouting,
		&analytics.Physiotherapy,
		&analytics.TotalSalaries,
		&analytics.TotalTraining,
		&analytics.TotalFinances,
		&analytics.TotalScouting,
		&analytics.TotalPhysiotherapy,
		&analytics.TrainingStaffCount,
		&analytics.FinanceStaffCount,
		&analytics.ScoutingStaffCount,
		&analytics.PhysiotherapyStaffCount,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No se encontraron datos GetAnalytics")
			return app.Analytics{}, nil
		}
		log.Printf("Error al escanear la fila: %v", err)
		return app.Analytics{}, err
	}

	return analytics, nil
}
