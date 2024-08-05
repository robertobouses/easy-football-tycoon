package lineup

import (
	"database/sql"
)

func (r *repository) PlayerExistsInLineup(playerId string) (bool, error) {
	var exists bool
	err := r.playerExistsInLineup.QueryRow(playerId).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, nil
}
