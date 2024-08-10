package continue

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetContinue() ([]app.Prospect, error)
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
