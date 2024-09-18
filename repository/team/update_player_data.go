package team

import (
	"log"

	"github.com/google/uuid"
)

func (r *repository) UpdatePlayerData(
	playerID uuid.UUID,
	playerName *string,
	position *string,
	age *int,
	fee *int,
	salary *int,
	technique *int,
	mental *int,
	physique *int,
	injuryDays *int,
	lined *bool,
	familiarity *int,
	fitness *int,
	happiness *int,
) error {
	log.Printf("Updating player data: %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v",
		playerID, playerName, position, age, fee, salary, technique, mental, physique, injuryDays, lined, familiarity, fitness, happiness)

	_, err := r.updatePlayerData.Exec(
		playerName,
		position,
		age,
		fee,
		salary,
		technique,
		mental,
		physique,
		injuryDays,
		lined,
		familiarity,
		fitness,
		happiness,
		playerID,
	)
	if err != nil {
		log.Printf("Error en updatePlayerLinedStatus repo: %v", err)
		return err
	}
	log.Println("Datos del jugador actualizados correctamente")
	return nil
}
