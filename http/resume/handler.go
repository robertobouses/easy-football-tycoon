package resume

import (
	"github.com/robertobouses/easy-football-tycoon/app"
	"github.com/robertobouses/easy-football-tycoon/app/staff"
	"github.com/robertobouses/easy-football-tycoon/http/signings"
)

type App interface {
	GetResume() ([]app.Calendary, error)

	ProcessPlayerSale() error
	SetCurrentPlayerSale(*app.Player, *int)
	GetCurrentPlayerSale() (*app.Player, *int, error)
	AcceptPlayerSale(player app.Player) error
	RejectPlayerSale(player app.Player)

	ProcessPlayerSigning() (signings.Signings, error)
	SetCurrentPlayerSigning(signings *signings.Signings)
	GetCurrentPlayerSigning() (*signings.Signings, error)
	AcceptPlayerSigning(signings *signings.Signings) error
	RejectPlayerSigning(signings *signings.Signings)

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

	ProcessInjury() (app.Player, error)
	GetCurrentInjuredPlayer() (*app.Player, *int, error)
	SetCurrentInjuredPlayer(*app.Player, *int)

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
