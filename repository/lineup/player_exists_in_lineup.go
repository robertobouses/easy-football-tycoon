package lineup

import (
	"database/sql"

	"github.com/google/uuid"
)

func (r *repository) PlayerExistsInLineup(playerId uuid.UUID) (bool, error) {
	var exists bool
	err := r.playerExistsInLineup.QueryRow(playerId).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, nil
}
