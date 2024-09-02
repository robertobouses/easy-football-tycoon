package resume

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetResume() ([]app.Calendary, error)
	ProcessSale() error
	AcceptSale(player app.Team) error
	RejectSale(player app.Team)
	ProcessInjury()
	ProcessMatch()
	GetCurrentSalePlayer() (*app.Team, error)
	SetCurrentSalePlayer(player *app.Team)
	GetCurrentProspect() (*app.Prospect, error)
	SetCurrentProspect(prospect *app.Prospect)
	AcceptPurchase(prospect *app.Prospect) error
	RejectPurchase(prospect *app.Prospect)
}

func NewHandler(app App) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
