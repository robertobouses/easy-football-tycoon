package app

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/robertobouses/easy-football-tycoon/app/staff"
)

func (a *AppService) ProcessTeamStaffSale() error {
	staffs, err := a.GetRandomStaff()
	if err != nil {
		return err
	}
	if staffs == (staff.Staff{}) {
		return errors.New("no staff available for sale")
	}

	transferFeeReceived, err := a.ProcessTransferFeeReceived(staffs.Fee)
	if err != nil {
		return err
	}

	a.SetCurrentStaffSale(&staffs, &transferFeeReceived)

	return nil
}

func (a *AppService) GetCurrentStaffSale() (*staff.Staff, *int, error) {
	return a.currentStaffOnSale, a.transferFeeReceived, nil
}

func (a *AppService) SetCurrentStaffSale(staff *staff.Staff, transferFeeReceived *int) {
	a.currentStaffOnSale = staff
	a.transferFeeReceived = transferFeeReceived
}

func (a *AppService) AcceptStaffSale(staff staff.Staff) error {
	if a.transferFeeReceived == nil {
		return ErrTransferNotFound
	}

	balance, err := a.bankRepo.GetBalance()
	if err != nil {
		return ErrBalanceNotFound
	}

	amount := *a.transferFeeReceived
	newBalance := balance + amount
	log.Println("amount en APP", amount)
	log.Println("balance inicial en APP", balance)
	log.Println("new balance en APP", newBalance)

	err = a.bankRepo.PostTransactions(amount, newBalance, staff.StaffName, "sale")
	if err != nil {
		return err
	}

	err = a.teamStaffRepo.DeleteTeamStaff(staff)
	if err != nil {
		return err
	}

	return nil
}

func (a *AppService) RejectStaffSale(staff staff.Staff) {
	a.SetCurrentStaffSale(nil, nil)
}

// TODO ROBERTO CREO QUE ESTE MÉTODO DE TRAER Y LUEGO HACER ALEATORIO ES MÁS ÓPTIMO QUE EL QUE LO HACE TODO EN REPO GetSigningsRandomByAnalytics
func (a *AppService) GetRandomStaff() (staff.Staff, error) {
	staffs, err := a.teamStaffRepo.GetTeamStaff()
	if err != nil {
		log.Println("Error al extraer GetStaff:", err)
		return staff.Staff{}, err
	}

	if len(staffs) == 0 {
		log.Println("No se encontraron jugadores")
		return staff.Staff{}, nil
	}

	log.Printf("Jugadores obtenidos: %+v", staffs)

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(staffs))
	staff := staffs[randomIndex]

	log.Printf("Empleado seleccionado aleatoriamente para la venta: %+v", staff)

	return staff, nil
}
