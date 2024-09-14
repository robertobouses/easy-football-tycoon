package match

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_matches.sql
var getMatchesQuery string

//go:embed sql/post_match.sql
var postMatchQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	postMatchStmt, err := db.Prepare(postMatchQuery)
	if err != nil {
		return nil, err
	}
	getMatchesStmt, err := db.Prepare(getMatchesQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:         db,
		postMatch:  postMatchStmt,
		getMatches: getMatchesStmt,
	}, nil
}

type repository struct {
	db         *sql.DB
	postMatch  *sql.Stmt
	getMatches *sql.Stmt
}
