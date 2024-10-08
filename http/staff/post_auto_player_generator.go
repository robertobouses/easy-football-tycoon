package staff

import (
	"log"
	"net/http"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

type PostAutoStaffGeneratorRequest struct {
	NumberOfStaffsToGenerate int `json:"numberofstafftogenerate"`
}

func (h Handler) PostAutoStaffGenerator(c *gin.Context) {
	var req PostAutoStaffGeneratorRequest
	if err := c.BindJSON(&req); err != nil {
		log.Printf("[PostAutoStaffGeneratorRequest] error parsing request: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.Printf("Número de empleados a generar: %d", req.NumberOfStaffsToGenerate)

	staffs, err := h.app.RunAutoStaffGenerator(req.NumberOfStaffsToGenerate)
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{
		"mensaje":             "Empleados generados correctamente",
		"empleados generados": staffs,
	})
}
