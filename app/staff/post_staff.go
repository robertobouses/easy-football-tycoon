package staff

import (
	"log"
)

func (a StaffService) PostStaff(req Staff) error {

	log.Printf("Valores de player creado: %+v\n", req)
	err := a.staffRepo.PostStaff(req)
	if err != nil {
		log.Println("Error en PostStaff:", err)
		return err
	}
	return nil
}
