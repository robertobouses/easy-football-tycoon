package staff

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/robertobouses/easy-football-tycoon/app"
)

type PostStaffRequest struct {
	StaffId       uuid.UUID `json:"staffid"`
	StaffName     string    `json:"staffname"`
	Job           string    `json:"job"`
	Age           int       `json:"age"`
	Fee           int       `json:"fee"`
	Salary        int       `json:"salary"`
	Training      int       `json:"trainin"`
	Finances      int       `json:"finances"`
	Scouting      int       `json:"scouting"`
	Physiotherapy int       `json:"physiotherapy"`
	Rarity        int       `json:"rarity"`
}

func (h Handler) PostStaff(c *gin.Context) {
	var req PostStaffRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostStaff] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	staff := app.Staff{
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

	err := h.app.PostStaff(staff)
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{"mensaje": "staff insertado correctamente"})
}
