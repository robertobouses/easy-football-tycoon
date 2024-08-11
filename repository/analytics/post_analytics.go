package analytics

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) PostAnalytics(req app.Analytics) error {
	_, err := r.postTeam.Exec(
		req.Finances,
		req.Scouting,
		req.Physiotherapy,
	)

	if err != nil {
		log.Print("Error en PostAnalytics repo", err)
		return err
	}

	return nil
}
