package team_staff

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

// TODO PODRÍA EVITAR DUCPLICAR ESTA FUNCIÓN Y LA DEL package staff PostStaff?
type PostTeamStaffRequest struct {
	StaffId       uuid.UUID `json:"staffid"`
	StaffName     string    `json:"staffname"`
	Job           string    `json:"job"`
	Age           int       `json:"age"`
	Fee           int       `json:"fee"`
	Salary        int       `json:"salary"`
	Training      int       `json:"training"`
	Finances      int       `json:"finances"`
	Scouting      int       `json:"scouting"`
	Physiotherapy int       `json:"physiotherapy"`
	Rarity        int       `json:"rarity"`
}

func (h Handler) PostTeamStaff(c *gin.Context) {
	var req PostTeamStaffRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostTeam] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	staff := staff.Staff{
		StaffId:       req.StaffId,
		StaffName:     req.StaffName,
		Job:           req.Job,
		Age:           req.Age,
		Fee:           req.Fee,
		Salary:        req.Salary,
		Training:      req.Training,
		Finances:      req.Finances,
		Scouting:      req.Scouting,
		Physiotherapy: req.Physiotherapy,
		Rarity:        req.Rarity,
	}
	err := h.app.PostTeamStaff(staff)

	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error al llamar desde http la app": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "jugador insertado correctamente"})
}
