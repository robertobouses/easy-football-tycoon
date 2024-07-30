package app

type Repository interface {
	GetLineup() ([]Lineup, error)
	GetTeam() ([]Team, error)
}

func NewApp(repo1, repo2 Repository) AppService {
	return AppService{
		lineupRepo: repo1,
		teamRepo:   repo2,
	}
}

type AppService struct {
	lineupRepo Repository
	teamRepo   Repository
}
