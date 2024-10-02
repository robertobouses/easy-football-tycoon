package staff

import (
	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

type App interface {
	PostStaff(staff.Staff) error
}

func NewHandler(app staff.StaffService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
