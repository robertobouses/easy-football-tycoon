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
	DeletePlayerFromTeam(player Team) error
}

type RivalRepository interface {
	GetRival() ([]Rival, error)
	PostRival(req Rival) error
}

type ProspectRepository interface {
	GetProspect() ([]Prospect, error)
	PostProspect(req Prospect) error
	GetProspectRandomByAnalytics(scouting int) (Prospect, error)
}

type CalendaryRepository interface {
	GetCalendary() ([]Calendary, error)
	PostCalendary(calendary []string) error
}

type AnalyticsRepository interface {
	GetAnalytics() (Analytics, error)
	PostAnalytics(req Analytics) error
}

func NewApp(lineupRepository LineupRepository, teamRepository TeamRepository, rivalRepository RivalRepository, prospectRepository ProspectRepository, calendaryRepository CalendaryRepository, analyticsRepository AnalyticsRepository) AppService {
	return AppService{
		lineupRepo:    lineupRepository,
		teamRepo:      teamRepository,
		rivalRepo:     rivalRepository,
		prospectRepo:  prospectRepository,
		calendaryRepo: calendaryRepository,
		analyticsRepo: analyticsRepository,
	}
}

type AppService struct {
	lineupRepo        LineupRepository
	teamRepo          TeamRepository
	rivalRepo         RivalRepository
	prospectRepo      ProspectRepository
	calendaryRepo     CalendaryRepository
	analyticsRepo     AnalyticsRepository
	currentSalePlayer *Team
	currentProspect   *Prospect
}
