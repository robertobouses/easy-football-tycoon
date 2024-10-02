package app

import (
	"github.com/google/uuid"
	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

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
		familiarity *int,
		fitness *int,
		happiness *int,
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
	DeleteSigning(signing Signings) error
}

type StaffRepository interface {
	GetStaff() ([]staff.Staff, error)
	GetStaffRandomByAnalytics(scouting int) (staff.Staff, error)
	DeleteStaffSigning(staff staff.Staff) error
}

type TeamStaffRepository interface {
	PostTeamStaff(req staff.Staff) error
	GetTeamStaff() ([]staff.Staff, error)
	DeleteTeamStaff(staff staff.Staff) error
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

type MatchRepository interface {
	PostMatch(match Match) error
	GetMatches() ([]Match, error)
}

type StrategyRepository interface {
	PostStrategy(Strategy) error
	GetStrategy() (Strategy, error)
}

func NewApp(lineupRepository LineupRepository, teamRepository TeamRepository, rivalRepository RivalRepository, signingsRepository SigningsRepository,
	staffRepository StaffRepository, teamStaffRepository TeamStaffRepository, calendaryRepository CalendaryRepository, analyticsRepository AnalyticsRepository,
	bankRepository BankRepository, matchRepository MatchRepository, strategyRepository StrategyRepository) AppService {
	return AppService{
		lineupRepo:    lineupRepository,
		teamRepo:      teamRepository,
		rivalRepo:     rivalRepository,
		signingsRepo:  signingsRepository,
		staffRepo:     staffRepository,
		teamStaffRepo: teamStaffRepository,
		calendaryRepo: calendaryRepository,
		analyticsRepo: analyticsRepository,
		bankRepo:      bankRepository,
		matchRepo:     matchRepository,
		strategyRepo:  strategyRepository,

		currentPlayerOnSale:  nil,
		currentPlayerSigning: nil,
		currentInjuredPlayer: nil,
		currentStaffSigning:  nil,
		currentStaffOnSale:   nil,
		injuryDays:           nil,
		transferFeeReceived:  nil,
		transferFeeIssued:    nil,
		callCounter:          1,
		callCounterRival:     1,
		currentRivals:        nil,
		isHome:               true,
		currentMatch:         nil,
	}
}

type AppService struct {
	lineupRepo    LineupRepository
	teamRepo      TeamRepository
	rivalRepo     RivalRepository
	signingsRepo  SigningsRepository
	staffRepo     StaffRepository
	teamStaffRepo TeamStaffRepository
	calendaryRepo CalendaryRepository
	analyticsRepo AnalyticsRepository
	bankRepo      BankRepository
	matchRepo     MatchRepository
	strategyRepo  StrategyRepository

	currentPlayerOnSale  *Team
	currentPlayerSigning *Signings
	currentInjuredPlayer *Team
	currentStaffSigning  *staff.Staff
	currentStaffOnSale   *staff.Staff
	injuryDays           *int
	transferFeeReceived  *int
	transferFeeIssued    *int

	callCounter      int
	callCounterRival int
	currentRivals    *[]Rival
	isHome           bool
	currentMatch     *Match
}
