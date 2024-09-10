package team_staff

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetTeamStaff() ([]app.Staff, error)
	PostTeamStaff(app.Staff) error
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
