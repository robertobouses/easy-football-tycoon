package resume

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetResume() ([]app.Calendary, error)
	ProcessSale() error
	AcceptSale(player app.Team) error
	RejectSale(player app.Team)
	ProcessMatch()
	GetCurrentSalePlayer() (*app.Team, *int, error)
	SetCurrentSalePlayer(*app.Team, *int)
	GetCurrentProspect() (*app.Prospect, error)
	SetCurrentProspect(prospect *app.Prospect)
	AcceptPurchase(prospect *app.Prospect) error
	RejectPurchase(prospect *app.Prospect)
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
