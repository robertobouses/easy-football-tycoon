package team_staff

import (
	"github.com/robertobouses/easy-football-tycoon/app"
	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

type App interface {
	GetTeamStaff() ([]staff.Staff, error)
	PostTeamStaff(staff.Staff) error
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
