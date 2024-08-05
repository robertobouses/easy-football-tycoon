package app

import (
	"log"
)

func (a AppService) PostLineup(req string) error {

	log.Printf("Valores de player creado: %+v\n", req)
	player, err := a.teamRepo.GetPlayerByPlayerId(req)
	if err != nil {
		log.Println("Error en GetPlayerByPlayerId:", err)
		return err
	}
	if player == nil {
		log.Printf("Jugador con playerid %s no encontrado", req)
		return ErrPlayerNotFound
	}
	exists, err := a.lineupRepo.PlayerExistsInLineup(req)
	if err != nil {
		log.Println("Error verificando si el jugador ya está alineado:", err)
		return err
	}
	if exists {
		return ErrPlayerAlreadyInLineup
	}

	err = a.lineupRepo.PostLineup(*player)
	if err != nil {
		log.Println("Error en PostLineup:", err)
		return err
	}
	err = a.teamRepo.UpdatePlayerLinedStatus(req)
	if err != nil {
		log.Println("Error en UpdatePlayerLinedStatus:", err)
		return err
	}
	return nil
}
