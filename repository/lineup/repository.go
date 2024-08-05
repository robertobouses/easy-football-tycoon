package lineup

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_lineup.sql
var getLineupQuery string

//go:embed sql/post_lineup.sql
var postLineupQuery string

//go:embed sql/player_exists_in_lineup.sql
var playerExistsInLineupQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getLineupStmt, err := db.Prepare(getLineupQuery)
	if err != nil {
		return nil, err
	}
	postLineupStmt, err := db.Prepare(postLineupQuery)
	if err != nil {
		return nil, err
	}
	playerExistsInLineupStmt, err := db.Prepare(playerExistsInLineupQuery)
	if err != nil {
		return nil, err
	}
	return &repository{
		db:                   db,
		getLineup:            getLineupStmt,
		postLineup:           postLineupStmt,
		playerExistsInLineup: playerExistsInLineupStmt,
	}, nil
}

type repository struct {
	db                   *sql.DB
	getLineup            *sql.Stmt
	postLineup           *sql.Stmt
	playerExistsInLineup *sql.Stmt
}
