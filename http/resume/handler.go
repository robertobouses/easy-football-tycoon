package resume

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetResume() ([]app.Prospect, error)
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
