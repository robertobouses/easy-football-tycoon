package match

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) PostMatch(match app.Match) error {
	_, err := r.postMatch.Exec(
		match.CalendaryId,
		match.MatchDayNumber,
		match.HomeOrAway,
		match.RivalName,
		match.Team.BallPossession,
		match.Team.ScoringChances,
		match.Team.Goals,
		match.Rival.BallPossession,
		match.Rival.ScoringChances,
		match.Rival.Goals,
	)

	if err != nil {
		log.Print("Error executing PostMatch statement:", err)
		return err
	}

	return nil
}
