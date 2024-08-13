package analytics

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetAnalytics() (app.Analytics, error)
	PostAnalytics(app.Analytics) error
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
