package team

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetTeam() ([]app.Team, error)
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
