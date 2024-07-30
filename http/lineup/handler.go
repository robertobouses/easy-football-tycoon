package lineup

import (
	"github.com/gin-gonic/gin"
	"github.com/robertobouses/easy-football-tycoon/app"
)

type App interface {
	GetLineup(ctx *gin.Context)
}

func NewHandler(app app.AppService) Handler {
	return Handler{
		App: app,
	}
}

type Handler struct {
	App App
}
