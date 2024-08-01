package team

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_team.sql
var getTeamQuery string

//go:embed sql/post_team.sql
var postTeamQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getTeamStmt, err := db.Prepare(getTeamQuery)
	if err != nil {
		return nil, err
	}
	postTeamStmt, err := db.Prepare(postTeamQuery)
	if err != nil {
		return nil, err
	}
	return &repository{
		db:       db,
		getTeam:  getTeamStmt,
		postTeam: postTeamStmt,
	}, nil
}

type repository struct {
	db       *sql.DB
	getTeam  *sql.Stmt
	postTeam *sql.Stmt
}
