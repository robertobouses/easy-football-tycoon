package resume

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetResume() ([]app.Calendary, error)
	ProcessPlayerSale() error
	AcceptPlayerSale(player app.Team) error
	RejectPlayerSale(player app.Team)

	ProcessMatch()

	GetCurrentSalePlayer() (*app.Team, *int, error)
	SetCurrentSalePlayer(*app.Team, *int)
	GetCurrentPlayerSigning() (*app.Signings, error)
	SetCurrentSigningPlayer(signings *app.Signings)

	AcceptPlayerSigning(signings *app.Signings) error
	RejectPlayerSigning(signings *app.Signings)

	ProcessInjury() (app.Team, error)

	GetCurrentInjuredPlayer() (*app.Team, *int, error)
	SetCurrentInjuredPlayer(*app.Team, *int)

	SetCurrentStaffSigning(staffSigning *app.Staff)
	GetCurrentStaffSigning() (*app.Staff, error)
	AcceptStaffSigning(staff *app.Staff) error
	RejectStaffSinging(staff *app.Staff)
	GetCurrentStaffSale() (*app.Staff, *int, error)
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
