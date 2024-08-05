package app

type LineupRepository interface {
	GetLineup() ([]Lineup, error)
	PostLineup(player Team) error
	PlayerExistsInLineup(playerId string) (bool, error)
}

type TeamRepository interface {
	GetTeam() ([]Team, error)
	PostTeam(req Team) error
	UpdatePlayerLinedStatus(req string) error
	GetPlayerByPlayerId(playerId string) (*Team, error)
}

func NewApp(repo1 LineupRepository, repo2 TeamRepository) AppService {
	return AppService{
		lineupRepo: repo1,
		teamRepo:   repo2,
	}
}

type AppService struct {
	lineupRepo LineupRepository
	teamRepo   TeamRepository
}
