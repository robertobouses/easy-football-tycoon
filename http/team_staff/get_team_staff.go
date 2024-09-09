package team_staff

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetTeamStaff(ctx *gin.Context) {
	staff, err := h.app.GetTeamStaff()
	if err != nil {
		ctx.JSON(nethttp.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(nethttp.StatusOK, staff)

}
