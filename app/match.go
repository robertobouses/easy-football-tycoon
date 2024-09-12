package app

import "github.com/google/uuid"

type TeamStats struct {
	BallPossession int `json:"ballpossession"`
	ScoringChances int `json:"scoringchances"`
	Goals          int `json:"goals"`
}

type Match struct {
	MatchId     uuid.UUID `json:"matchid"`
	CalendaryId int       `json:"calendaryid"`
	RivalName   string    `json:"rivalname"`
	Team        TeamStats `json:"team"`
	Rival       TeamStats `json:"rival"`
}

//TODO vistosidad?? y que sea motivo de despido
