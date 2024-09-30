package match

import (
	_ "github.com/lib/pq"
	"github.com/robertobouses/easy-football-tycoon/app"
)

func (r *repository) GetMatches() ([]app.Match, error) {
	rows, err := r.getMatches.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var matches []app.Match

	for rows.Next() {
		var match app.Match
		var homeOrAway string

		err := rows.Scan(
			&match.MatchId,
			&match.CalendaryId,
			&match.MatchDayNumber,
			&homeOrAway,
			&match.RivalName,
			&match.Team.BallPossession,
			&match.Team.ScoringChances,
			&match.Team.Goals,
			&match.Rival.BallPossession,
			&match.Rival.ScoringChances,
			&match.Rival.Goals,
			&match.Result,
		)
		if err != nil {
			return nil, err
		}

		// match.HomeOrAway = HomeOrAway(homeOrAway)
		matches = append(matches, match)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}
