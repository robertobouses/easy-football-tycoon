package lineup

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_lineup.sql
var getLineupQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getLineupStmt, err := db.Prepare(getLineupQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:        db,
		getLineup: getLineupStmt,
	}, nil
}
