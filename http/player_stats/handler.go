package stats

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetPlayerStats() ([]app.PlayerStats, error)
}

func NewHandler(app *app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
