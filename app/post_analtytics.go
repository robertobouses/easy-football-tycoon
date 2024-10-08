package app

import (
	"fmt"
	"log"
)

func (a AppService) PostAnalytics(req Analytics, job string) error {

	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("Error obteniendo Analytics:", err)
		return err
	}

	req.TotalTraining = analytics.TotalTraining + req.Training
	req.TotalFinances = analytics.TotalFinances + req.Finances
	req.TotalScouting = analytics.TotalScouting + req.Scouting
	req.TotalPhysiotherapy = analytics.TotalPhysiotherapy + req.Physiotherapy

	switch job {
	case "trainer":
		req.TrainingStaffCount = analytics.TrainingStaffCount + 1
		req.FinanceStaffCount = analytics.FinanceStaffCount
		req.ScoutingStaffCount = analytics.ScoutingStaffCount
		req.PhysiotherapyStaffCount = analytics.PhysiotherapyStaffCount
	case "financial":
		req.TrainingStaffCount = analytics.TrainingStaffCount
		req.FinanceStaffCount = analytics.FinanceStaffCount + 1
		req.ScoutingStaffCount = analytics.ScoutingStaffCount
		req.PhysiotherapyStaffCount = analytics.PhysiotherapyStaffCount
	case "scout":
		req.TrainingStaffCount = analytics.TrainingStaffCount
		req.FinanceStaffCount = analytics.FinanceStaffCount
		req.ScoutingStaffCount = analytics.ScoutingStaffCount + 1
		req.PhysiotherapyStaffCount = analytics.PhysiotherapyStaffCount
	case "physiotherapist":
		req.TrainingStaffCount = analytics.TrainingStaffCount
		req.FinanceStaffCount = analytics.FinanceStaffCount
		req.ScoutingStaffCount = analytics.ScoutingStaffCount
		req.PhysiotherapyStaffCount = analytics.PhysiotherapyStaffCount + 1
	case "movimientos iniciales":
		req.TrainingStaffCount = analytics.TrainingStaffCount
		req.FinanceStaffCount = analytics.FinanceStaffCount
		req.ScoutingStaffCount = analytics.ScoutingStaffCount
		req.PhysiotherapyStaffCount = analytics.PhysiotherapyStaffCount
	default:
		log.Printf("Tipo de trabajo no reconocido: %s", job)
		return fmt.Errorf("tipo de trabajo no reconocido: %s", job)
	}

	log.Printf("Valores de Analytics antes de enviar: %+v\n", req)

	err = a.analyticsRepo.PostAnalytics(req, job)
	if err != nil {
		log.Println("Error en PostAnalytics:", err)
		return err
	}

	return nil
}
