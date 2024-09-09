package team_staff

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_team_staff.sql
var getTeamStaffQuery string

//go:embed sql/post_team_staff.sql
var postTeamStaffQuery string

//go:embed sql/delete_team_staff.sql
var deleteTeamStaffQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getTeamStaffStmt, err := db.Prepare(getTeamStaffQuery)
	if err != nil {
		return nil, err
	}
	postTeamStaffStmt, err := db.Prepare(postTeamStaffQuery)
	if err != nil {
		return nil, err
	}

	deleteTeamStaffStmt, err := db.Prepare(deleteTeamStaffQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:              db,
		getTeamStaff:    getTeamStaffStmt,
		postTeamStaff:   postTeamStaffStmt,
		deleteTeamStaff: deleteTeamStaffStmt,
	}, nil
}

type repository struct {
	db              *sql.DB
	getTeamStaff    *sql.Stmt
	postTeamStaff   *sql.Stmt
	deleteTeamStaff *sql.Stmt
}
