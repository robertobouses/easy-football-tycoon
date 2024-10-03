package team

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetTeam() ([]app.Player, error)
	PostTeam(app.Player) error
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
