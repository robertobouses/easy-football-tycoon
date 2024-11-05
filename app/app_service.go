package app

import (
	"github.com/google/uuid"
	"github.com/robertobouses/easy-football-tycoon/app/signings"
	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

type LineupRepository interface {
	GetLineup() ([]Lineup, error)
	PostLineup(player Player) error
	PlayerExistsInLineup(playerId uuid.UUID) (bool, error)
	DeletePlayerFromLineup(playerId uuid.UUID) error
}

type TeamRepository interface {
	GetTeam() ([]Player, error)
	PostTeam(req Player) error
	UpdatePlayerLinedStatus(req uuid.UUID) error
	GetPlayerByPlayerId(playerId uuid.UUID) (*Player, error)
	DeletePlayerFromTeam(player Player) error
	UpdatePlayerData(
		playerID uuid.UUID,
		firstName *string,
		lastName *string,
		nationality *string,
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
	GetSignings() ([]signings.Signings, error)
	PostSignings(req signings.Signings) error
	GetSigningsRandomByAnalytics(scouting int) (signings.Signings, error)
	DeleteSigning(signing signings.Signings) error
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

type StatsRepository interface {
	UpdatePlayerStats(playerId uuid.UUID, appearances int, blocks int, saves int, aerialDuelsWon int, keyPass int, assists int, chances int, goals int, mvp int, rating float64) error
	GetPlayerStats() ([]PlayerStats, error)
	GetPlayerStatsById(playerId uuid.UUID) (*PlayerStats, error)
}

func NewApp(
	lineupRepository LineupRepository,
	teamRepository TeamRepository,
	rivalRepository RivalRepository,
	signingsRepository SigningsRepository,
	staffRepository StaffRepository,
	teamStaffRepository TeamStaffRepository,
	calendaryRepository CalendaryRepository,
	analyticsRepository AnalyticsRepository,
	bankRepository BankRepository,
	matchRepository MatchRepository,
	strategyRepository StrategyRepository,
	statsRepository StatsRepository) AppService {
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
		statsRepo:     statsRepository,

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

		currentePlayerOnMatchEvent: nil,
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
	statsRepo     StatsRepository

	currentPlayerOnSale  *Player
	currentPlayerSigning *signings.Signings
	currentInjuredPlayer *Player
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

	currentePlayerOnMatchEvent *Lineup
}
