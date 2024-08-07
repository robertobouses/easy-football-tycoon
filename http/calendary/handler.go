package calendary

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetCalendary() ([]app.Calendary, error)
	PostCalendary() error
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
