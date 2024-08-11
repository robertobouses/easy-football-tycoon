package analytics

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_prospect.sql
var getProspectQuery string

//go:embed sql/post_prospect.sql
var postProspectQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getProspectStmt, err := db.Prepare(getProspectQuery)
	if err != nil {
		return nil, err
	}
	postProspectStmt, err := db.Prepare(postProspectQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:           db,
		getProspect:  getProspectStmt,
		postProspect: postProspectStmt,
	}, nil
}

type repository struct {
	db           *sql.DB
	getProspect  *sql.Stmt
	postProspect *sql.Stmt
}
