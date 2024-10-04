package signings

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app/signings"
)

func (r *repository) DeleteSigning(signing signings.Signings) error {
	log.Printf("Intentando eliminar jugador con signingId: %v", signing.SigningsId)
	_, err := r.deleteSigning.Exec(signing.SigningsId)
	if err != nil {
		log.Printf("Error en deletesigningFromTeam repo: %v", err)
		return err
	}
	log.Printf("El jugador %v se fue de tu equipo", signing.SigningsId)
	return nil
}
