package team

import (
	"log"
)

func (r *repository) UpdatePlayerLinedStatus(playerID string) error {
	_, err := r.updatePlayerLinedStatus.Exec(playerID)
	if err != nil {
		log.Printf("Error en updatePlayerLinedStatus repo: %v", err)
		return err
	}
	log.Println("Estado de 'lined' actualizado correctamente")
	return nil
}
