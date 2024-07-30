package lineup

import (
	"log"

	"github.com/robertobouses/easy-football-tycoon/app"
	"github.com/robertobouses/easy-football-tycoon/repository" // Importa el paquete repository
)

type Handler struct {
	repo *repository.Repository // Asegúrate de usar el tipo correcto
}

func (h *Handler) GetLineup() ([]app.Lineup, error) {
	rows, err := h.repo.getLineup.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lineups []app.Lineup
	for rows.Next() {
		var lineup app.Lineup
		if err := rows.Scan(
			&lineup.PlayerName,
			&lineup.Position,
			&lineup.Technique,
			&lineup.Mental,
			&lineup.Physique,
		); err != nil {
			log.Printf("Error al escanear las filas: %v", err)
			return nil, err
		}
		lineups = append(lineups, lineup)
	}

	return lineups, nil
}
