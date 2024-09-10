package signings

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetSignings() ([]app.Signings, error)
	PostSignings(app.Signings) error
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
