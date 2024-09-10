package signings

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_signings.sql
var getSigningsQuery string

//go:embed sql/post_signings.sql
var postSigningsQuery string

//go:embed sql/get_signings_random_by_analytics.sql
var getSigningsRandomByAnalyticsQuery string

//go:embed sql/delete_signing.sql
var deleteSigningQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getSigningsStmt, err := db.Prepare(getSigningsQuery)
	if err != nil {
		return nil, err
	}
	postSigningsStmt, err := db.Prepare(postSigningsQuery)
	if err != nil {
		return nil, err
	}
	getSigningsRandomByAnalyticsStmt, err := db.Prepare(getSigningsRandomByAnalyticsQuery)
	if err != nil {
		return nil, err
	}
	deleteSigningStmt, err := db.Prepare(deleteSigningQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:                           db,
		getSignings:                  getSigningsStmt,
		postSignings:                 postSigningsStmt,
		getSigningsRandomByAnalytics: getSigningsRandomByAnalyticsStmt,
		deleteSigning:                deleteSigningStmt,
	}, nil
}

type repository struct {
	db                           *sql.DB
	getSignings                  *sql.Stmt
	postSignings                 *sql.Stmt
	getSigningsRandomByAnalytics *sql.Stmt
	deleteSigning                *sql.Stmt
}
