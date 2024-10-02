package resume

import (
	"github.com/robertobouses/easy-football-tycoon/app"
	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

type App interface {
	GetResume() ([]app.Calendary, error)

	ProcessPlayerSale() error
	SetCurrentPlayerSale(*app.Team, *int)
	GetCurrentPlayerSale() (*app.Team, *int, error)
	AcceptPlayerSale(player app.Team) error
	RejectPlayerSale(player app.Team)

	ProcessPlayerSigning() (app.Signings, error)
	SetCurrentPlayerSigning(signings *app.Signings)
	GetCurrentPlayerSigning() (*app.Signings, error)
	AcceptPlayerSigning(signings *app.Signings) error
	RejectPlayerSigning(signings *app.Signings)

	ProcessStaffSigning() (staff.Staff, error)
	SetCurrentStaffSigning(staffSigning *staff.Staff)
	GetCurrentStaffSigning() (*staff.Staff, error)
	AcceptStaffSigning(staff *staff.Staff) error
	RejectStaffSigning(staff *staff.Staff)

	ProcessTeamStaffSale() error
	SetCurrentStaffSale(*staff.Staff, *int)
	GetCurrentStaffSale() (*staff.Staff, *int, error)
	AcceptStaffSale(staff staff.Staff) error
	RejectStaffSale(staff staff.Staff)

	ProcessMatch(int) (app.Match, error)

	ProcessInjury() (app.Team, error)
	GetCurrentInjuredPlayer() (*app.Team, *int, error)
	SetCurrentInjuredPlayer(*app.Team, *int)

	GetCurrentRival() (app.Rival, string, int, error)
	SetCurrentMatch(match *app.Match)
	GetCurrentMatch() (*app.Match, error)
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
