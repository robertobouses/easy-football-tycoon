package app

import (
	"log"

	"github.com/google/uuid"
)

func (a AppService) PostLineup(playerid uuid.UUID) error {

	log.Printf("Valores de player creado: %+v\n", playerid)
	player, err := a.teamRepo.GetPlayerByPlayerId(playerid)
	if err != nil {
		log.Println("Error en GetPlayerByPlayerId:", err)
		return err
	}
	if player == nil {
		log.Printf("Jugador con playerid %s no encontrado", playerid)
		return ErrPlayerNotFound
	}
	exists, err := a.lineupRepo.PlayerExistsInLineup(playerid)
	if err != nil {
		log.Println("Error verificando si el jugador ya está alineado:", err)
		return err
	}
	if exists {
		return ErrPlayerAlreadyInLineup
	}

	lineup, err := a.lineupRepo.GetLineup()
	if err != nil {
		log.Println("Error obteniendo la alineación:", err)
		return err
	}
	if len(lineup) >= 11 {
		return ErrLineupCompleted
	}

	err = a.lineupRepo.PostLineup(*player)
	if err != nil {
		log.Println("Error en PostLineup:", err)
		return err
	}
	err = a.teamRepo.UpdatePlayerLinedStatus(playerid)
	if err != nil {
		log.Println("Error en UpdatePlayerLinedStatus:", err)
		return err
	}
	return nil
}
