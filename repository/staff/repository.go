package staff

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/post_staff.sql
var postStaffQuery string

func NewRepository(db *sql.DB) (*repository, error) {

	postStaffStmt, err := db.Prepare(postStaffQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:        db,
		postStaff: postStaffStmt,
	}, nil
}

type repository struct {
	db        *sql.DB
	postStaff *sql.Stmt
}
