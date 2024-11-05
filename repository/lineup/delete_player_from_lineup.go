package lineup

import (
	"log"

	"github.com/google/uuid"
)

func (r *repository) DeletePlayerFromLineup(playerId uuid.UUID) error {
	log.Printf("Intentando eliminar jugador de la alineación signingId: %v", playerId)
	_, err := r.deletePlayerFromLineup.Exec(playerId)
	if err != nil {
		log.Printf("Error en DeletePlayerFromLineup repo: %v", err)
		return err
	}
	log.Printf("El jugador %v salió de la alineación", playerId)
	return nil
}
