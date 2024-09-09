package app

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

func (a *AppService) ProcessStaffSale() error {
	staff, err := a.GetRandomStaff()
	if err != nil {
		return err
	}
	if staff == (Staff{}) {
		return errors.New("no staff available for sale")
	}

	transferFeeReceived, err := a.ProcessTransferFeeReceived(staff.Fee)
	if err != nil {
		return err
	}

	a.SetCurrentSaleStaff(&staff, &transferFeeReceived)

	return nil
}

func (a *AppService) GetCurrentStaffSale() (*Staff, *int, error) {
	return a.currentStaffOnSale, a.transferFeeReceived, nil
}

func (a *AppService) SetCurrentSaleStaff(staff *Staff, transferFeeReceived *int) {
	a.currentStaffOnSale = staff
	a.transferFeeReceived = transferFeeReceived
}

func (a *AppService) AcceptStaffSale(staff Staff) error {
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

func (a *AppService) RejectStaffSale(staff Staff) {
	a.SetCurrentSaleStaff(nil, nil)
}

// TODO ROBERTO CREO QUE ESTE MÉTODO DE TRAER Y LUEGO HACER ALEATORIO ES MÁS ÓPTIMO QUE EL QUE LO HACE TODO EN REPO GetSigningsRandomByAnalytics
func (a *AppService) GetRandomStaff() (Staff, error) {
	staffs, err := a.staffRepo.GetStaff()
	if err != nil {
		log.Println("Error al extraer GetStaff:", err)
		return Staff{}, err
	}

	if len(staffs) == 0 {
		log.Println("No se encontraron jugadores")
		return Staff{}, nil
	}

	log.Printf("Jugadores obtenidos: %+v", staffs)

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(staffs))
	staff := staffs[randomIndex]

	log.Printf("Jugador seleccionado aleatoriamente para la venta: %+v", staff)

	return staff, nil
}
