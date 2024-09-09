package app

import (
	"log"

	"github.com/google/uuid"
)

func (a *AppService) ProcessStaffSigning() (Staff, error) {
	analytics, err := a.analyticsRepo.GetAnalytics()
	if err != nil {
		log.Println("Error al extraer GetAnalytics", err)
		return Staff{}, err
	}

	staffSigning, err := a.staffRepo.GetStaffRandomByAnalytics(analytics.Scouting)
	log.Println("Signingso asignado en ProcessStaffSigning 1:", staffSigning)
	if err != nil {
		log.Println("Error al extraer GetSigningsRandomByAnalytics", err)
		return Staff{}, err
	}

	if staffSigning.StaffId == uuid.Nil {
		log.Println("No se encontró un signingso válido")
		return Staff{}, nil
	}

	a.SetCurrentStaffSigning(&staffSigning)
	log.Println("Signingso asignado en ProcessStaffSigning 2:", staffSigning)

	return staffSigning, nil
}

func (a *AppService) SetCurrentStaffSigning(staffSigning *Staff) {
	if staffSigning == nil || staffSigning.StaffId == uuid.Nil {
		log.Println("Signingso no válido, no se asignará. Signingso:", staffSigning)
		a.currentStaffSigning = nil
	} else {
		log.Println("Signingso asignado en SetcurrentStaffSigningtaff:", *staffSigning)
		a.currentStaffSigning = staffSigning
	}
}

func (a *AppService) GetCurrentStaffSigning() (*Staff, error) {
	log.Println("a.currentStaffSigning en GetCurrentStaffSigning 1:", a.currentStaffSigning)
	if a.currentStaffSigning == nil {
		log.Println("a.currentStaffSigning es nil en GetCurrentStaffSigning 2")
		return nil, nil
	}
	log.Println("Signingso actual en GetCurrentStaffSigning 3:", *a.currentStaffSigning)
	return a.currentStaffSigning, nil
}

func (a *AppService) AcceptStaffSigning(signings *Staff) error {

	initialFee := signings.Fee

	paid, err := a.ProcessTransferFeePaid(initialFee)
	if err != nil {
		return ErrPaidNotFound
	}

	initalBalance, err := a.bankRepo.GetBalance()
	if err != nil {
		return ErrBalanceNotFound
	}

	newBalance := initalBalance - paid

	a.bankRepo.PostTransactions(paid, newBalance, signings.StaffName, "Staff Signing")

	a.teamStaffRepo.PostTeamStaff(*signings)
	return nil
}

func (a *AppService) RejectStaffSinging(signings *Staff) {
	a.SetCurrentStaffSigning(nil)
}
