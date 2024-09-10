package resume

import (
	"github.com/robertobouses/easy-football-tycoon/app"
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

	ProcessStaffSigning() (app.Staff, error)
	SetCurrentStaffSigning(staffSigning *app.Staff)
	GetCurrentStaffSigning() (*app.Staff, error)
	AcceptStaffSigning(staff *app.Staff) error
	RejectStaffSigning(staff *app.Staff)

	ProcessTeamStaffSale() error
	SetCurrentStaffSale(*app.Staff, *int)
	GetCurrentStaffSale() (*app.Staff, *int, error)
	AcceptStaffSale(staff app.Staff) error
	RejectStaffSale(staff app.Staff)

	ProcessMatch()

	ProcessInjury() (app.Team, error)
	GetCurrentInjuredPlayer() (*app.Team, *int, error)
	SetCurrentInjuredPlayer(*app.Team, *int)
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
