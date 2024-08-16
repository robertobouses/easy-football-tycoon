package team

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_team.sql
var getTeamQuery string

//go:embed sql/post_team.sql
var postTeamQuery string

//go:embed sql/update_player_lined_status.sql
var updatePlayerLinedStatusQuery string

//go:embed sql/get_player_by_playerid.sql
var getPlayerByPlayerIdQuery string

//go:embed sql/delete_player_from_team.sql
var deletePlayerFromTeamQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getTeamStmt, err := db.Prepare(getTeamQuery)
	if err != nil {
		return nil, err
	}
	postTeamStmt, err := db.Prepare(postTeamQuery)
	if err != nil {
		return nil, err
	}

	updatePlayerLinedStatusStmt, err := db.Prepare(updatePlayerLinedStatusQuery)
	if err != nil {
		return nil, err
	}

	getPlayerByPlayerIdStmt, err := db.Prepare(getPlayerByPlayerIdQuery)
	if err != nil {
		return nil, err
	}
	deletePlayerFromTeamStmt, err := db.Prepare(deletePlayerFromTeamQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:                      db,
		getTeam:                 getTeamStmt,
		postTeam:                postTeamStmt,
		updatePlayerLinedStatus: updatePlayerLinedStatusStmt,
		getPlayerByPlayerId:     getPlayerByPlayerIdStmt,
		deletePlayerFromTeam:    deletePlayerFromTeamStmt,
	}, nil
}

type repository struct {
	db                      *sql.DB
	getTeam                 *sql.Stmt
	postTeam                *sql.Stmt
	updatePlayerLinedStatus *sql.Stmt
	getPlayerByPlayerId     *sql.Stmt
	deletePlayerFromTeam    *sql.Stmt
}
