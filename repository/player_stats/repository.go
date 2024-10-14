package player_stats

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/update_player_stats.sql
var updatePlayerStatsQuery string

//go:embed sql/get_player_stats.sql
var getPlayerStatsQuery string

func NewRepository(db *sql.DB) (*repository, error) {

	updatePlayerStatsStmt, err := db.Prepare(updatePlayerStatsQuery)
	if err != nil {
		return nil, err
	}

	getPlayerStatsStmt, err := db.Prepare(getPlayerStatsQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:                db,
		updatePlayerStats: updatePlayerStatsStmt,
		getPlayerStats:    getPlayerStatsStmt,
	}, nil
}

type repository struct {
	db                *sql.DB
	updatePlayerStats *sql.Stmt
	getPlayerStats    *sql.Stmt
}
