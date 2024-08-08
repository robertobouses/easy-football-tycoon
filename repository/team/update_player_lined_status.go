package team

import (
	"log"

	"github.com/google/uuid"
)

func (r *repository) UpdatePlayerLinedStatus(playerID uuid.UUID) error {
	_, err := r.updatePlayerLinedStatus.Exec(playerID)
	if err != nil {
		log.Printf("Error en updatePlayerLinedStatus repo: %v", err)
		return err
	}
	log.Println("Estado de 'lined' actualizado correctamente")
	return nil
}
