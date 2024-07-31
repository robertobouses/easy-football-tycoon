package lineup

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetLineup() ([]app.Lineup, error)
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
