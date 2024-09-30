package analytics

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) PostAnalytics(req app.Analytics, job string) error {
	_, err := r.postAnalytics.Exec(
		req.Training,
		req.Finances,
		req.Scouting,
		req.Physiotherapy,
		req.TotalSalaries,
		req.TotalTraining,
		req.TotalFinances,
		req.TotalScouting,
		req.TotalPhysiotherapy,
		req.TrainingStaffCount,
		req.FinanceStaffCount,
		req.ScoutingStaffCount,
		req.PhysiotherapyStaffCount,
		req.Trust,
		req.StadiumCapacity,
		job,
	)

	if err != nil {
		log.Print("Error en PostAnalytics repo:", err)
		return err
	}

	return nil
}
