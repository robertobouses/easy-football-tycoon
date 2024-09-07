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
	UpdatePlayerData(
		playerID uuid.UUID,
		playerName *string,
		position *string,
		age *int,
		fee *int,
		salary *int,
		technique *int,
		mental *int,
		physique *int,
		injuryDays *int,
		lined *bool,
	) error
}

type RivalRepository interface {
	GetRival() ([]Rival, error)
	PostRival(req Rival) error
}

type SigningsRepository interface {
	GetSignings() ([]Signings, error)
	PostSignings(req Signings) error
	GetSigningsRandomByAnalytics(scouting int) (Signings, error)
}

type StaffRepository interface {
	PostStaff(req Staff) error
	//GetSigningsRandomByAnalytics(scouting int) (Signings, error)
}

type CalendaryRepository interface {
	GetCalendary() ([]Calendary, error)
	PostCalendary(calendary []string) error
}

type AnalyticsRepository interface {
	GetAnalytics() (Analytics, error)
	PostAnalytics(req Analytics, job string) error
}

type BankRepository interface {
	PostTransactions(amount int, balance int, signings string, transactionType string) error
	GetBalance() (int, error)
}

func NewApp(lineupRepository LineupRepository, teamRepository TeamRepository, rivalRepository RivalRepository, signingsRepository SigningsRepository, staffRepository StaffRepository, calendaryRepository CalendaryRepository, analyticsRepository AnalyticsRepository, bankRepository BankRepository) AppService {
	return AppService{
		lineupRepo:    lineupRepository,
		teamRepo:      teamRepository,
		rivalRepo:     rivalRepository,
		signingsRepo:  signingsRepository,
		staffRepo:     staffRepository,
		calendaryRepo: calendaryRepository,
		analyticsRepo: analyticsRepository,
		bankRepo:      bankRepository,
	}
}

type AppService struct {
	lineupRepo           LineupRepository
	teamRepo             TeamRepository
	rivalRepo            RivalRepository
	signingsRepo         SigningsRepository
	staffRepo            StaffRepository
	calendaryRepo        CalendaryRepository
	analyticsRepo        AnalyticsRepository
	bankRepo             BankRepository
	currentSalePlayer    *Team
	currentSignings      *Signings
	currentInjuredPlayer *Team
	injuryDays           *int
	transferFeeReceived  *int
	callCounter          int
}
