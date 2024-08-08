package lineup

import (
	"github.com/google/uuid"
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetLineup() ([]app.Lineup, error)
	PostLineup(req uuid.UUID) error
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
