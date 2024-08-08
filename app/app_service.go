package app

import "github.com/google/uuid"

type LineupRepository interface {
	GetLineup() ([]Lineup, error)
	PostLineup(player Team) error
	PlayerExistsInLineup(playerId uuid.UUID) (bool, error)
}

type TeamRepository interface {
	GetTeam() ([]Team, error)
	PostTeam(req Team) error
	UpdatePlayerLinedStatus(req uuid.UUID) error
	GetPlayerByPlayerId(playerId uuid.UUID) (*Team, error)
}

type RivalRepository interface {
	GetRival() ([]Rival, error)
	PostRival(req Rival) error
}

type ProspectRepository interface {
	GetProspect() ([]Prospect, error)
	PostProspect(req Prospect) error
}

type CalendaryRepository interface {
	GetCalendary() ([]Calendary, error)
	PostCalendary(calendary []string) error
}

func NewApp(repo1 LineupRepository, repo2 TeamRepository, repo3 RivalRepository, repo4 ProspectRepository, repo5 CalendaryRepository) AppService {
	return AppService{
		lineupRepo:    repo1,
		teamRepo:      repo2,
		rivalRepo:     repo3,
		prospectRepo:  repo4,
		calendaryRepo: repo5,
	}
}

type AppService struct {
	lineupRepo    LineupRepository
	teamRepo      TeamRepository
	rivalRepo     RivalRepository
	prospectRepo  ProspectRepository
	calendaryRepo CalendaryRepository
}
