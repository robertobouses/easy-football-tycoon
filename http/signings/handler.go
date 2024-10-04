package signings

import (
	"github.com/robertobouses/easy-football-tycoon/app/signings"
)

type App interface {
	GetSignings() ([]signings.Signings, error)
	PostSignings(signings.Signings) error
	RunAutoPlayerGenerator(numberOfPlayers int) ([]signings.Signings, error)
}

func NewHandler(app *signings.SigningsService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
