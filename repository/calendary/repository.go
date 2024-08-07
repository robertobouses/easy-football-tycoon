package calendary

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_calendary.sql
var getCalendaryQuery string

//go:embed sql/post_calendary.sql
var postCalendaryQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getCalendaryStmt, err := db.Prepare(getCalendaryQuery)
	if err != nil {
		return nil, err
	}
	postCalendaryStmt, err := db.Prepare(postCalendaryQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:            db,
		getCalendary:  getCalendaryStmt,
		postCalendary: postCalendaryStmt,
	}, nil
}

type repository struct {
	db            *sql.DB
	getCalendary  *sql.Stmt
	postCalendary *sql.Stmt
}
