package app

import "github.com/google/uuid"

type Match struct {
	MatchId        uuid.UUID  `json:"matchid"`
	CalendaryId    int        `json:"calendaryid"`
	MatchDayNumber int        `json:"matchdaynumber"`
	HomeOrAway     HomeOrAway `json:"homeoraway"`
	RivalName      string     `json:"rivalname"`
	Team           TeamStats  `json:"team"`
	Rival          TeamStats  `json:"rival"`
	Result         string     `json:"result"`
}

//TODO vistosidad?? y que sea motivo de despido o sirve con chances

type TeamStats struct {
	BallPossession int `json:"ballpossession"`
	ScoringChances int `json:"scoringchances"`
	Goals          int `json:"goals"`
}
type HomeOrAway string

const (
	Home HomeOrAway = "home"
	Away HomeOrAway = "away"
)
