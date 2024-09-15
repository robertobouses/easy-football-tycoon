package strategy

import (
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetStrategy() (app.Strategy, error)
	PostStrategy(app.Strategy) error
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		app: app,
	}
}

type Handler struct {
	app App
}
