package signings

type SigningsRepository interface {
	GetSignings() ([]Signings, error)
	PostSignings(req Signings) error
	GetSigningsRandomByAnalytics(scouting int) (Signings, error)
	DeleteSigning(signing Signings) error
}

func NewApp(signingsRepository SigningsRepository) SigningsService {
	return SigningsService{
		signingsRepo: signingsRepository,
	}
}

type SigningsService struct {
	signingsRepo SigningsRepository
}
