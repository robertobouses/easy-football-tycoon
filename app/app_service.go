package app

type LineupRepository interface {
	GetLineup() ([]Lineup, error)
}

type TeamRepository interface {
	GetTeam() ([]Team, error)
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
