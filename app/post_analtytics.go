package app

import (
	"fmt"
	"log"
)

func (a AppService) PostAnalytics(req Analytics, job string) error {
	// Obtener los datos actuales de Analytics
	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("Error obteniendo Analytics:", err)
		return err
	}

	// Actualizar los totales
	req.TotalTraining = analytics.TotalTraining + req.Training
	req.TotalFinances = analytics.TotalFinances + req.Finances
	req.TotalScouting = analytics.TotalScouting + req.Scouting
	req.TotalPhysiotherapy = analytics.TotalPhysiotherapy + req.Physiotherapy

	// Actualizar los contadores de personal basados en el tipo de trabajo, solo si es un trabajo reconocido
	switch job {
	case "trainer":
		req.TrainingStaffCount = analytics.TrainingStaffCount + 1
		req.FinanceStaffCount = analytics.FinanceStaffCount
		req.ScoutingStaffCount = analytics.ScoutingStaffCount
		req.PhysiotherapyStaffCount = analytics.PhysiotherapyStaffCount
	case "finances":
		req.TrainingStaffCount = analytics.TrainingStaffCount
		req.FinanceStaffCount = analytics.FinanceStaffCount + 1
		req.ScoutingStaffCount = analytics.ScoutingStaffCount
		req.PhysiotherapyStaffCount = analytics.PhysiotherapyStaffCount
	case "scouting":
		req.TrainingStaffCount = analytics.TrainingStaffCount
		req.FinanceStaffCount = analytics.FinanceStaffCount
		req.ScoutingStaffCount = analytics.ScoutingStaffCount + 1
		req.PhysiotherapyStaffCount = analytics.PhysiotherapyStaffCount
	case "physiotherapy":
		req.TrainingStaffCount = analytics.TrainingStaffCount
		req.FinanceStaffCount = analytics.FinanceStaffCount
		req.ScoutingStaffCount = analytics.ScoutingStaffCount
		req.PhysiotherapyStaffCount = analytics.PhysiotherapyStaffCount + 1
	case "movimientos iniciales":
		// No se actualizan contadores si el job es "movimientos iniciales"
		req.TrainingStaffCount = analytics.TrainingStaffCount
		req.FinanceStaffCount = analytics.FinanceStaffCount
		req.ScoutingStaffCount = analytics.ScoutingStaffCount
		req.PhysiotherapyStaffCount = analytics.PhysiotherapyStaffCount
	default:
		log.Printf("Tipo de trabajo no reconocido: %s", job)
		return fmt.Errorf("tipo de trabajo no reconocido: %s", job)
	}

	log.Printf("Valores de Analytics antes de enviar: %+v\n", req)

	// Insertar los nuevos datos en la base de datos
	err = a.analyticsRepo.PostAnalytics(req, job)
	if err != nil {
		log.Println("Error en PostAnalytics:", err)
		return err
	}

	return nil
}
