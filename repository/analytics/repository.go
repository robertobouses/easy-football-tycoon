package analytics

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_analytics.sql
var getAnalyticsQuery string

//go:embed sql/post_analytics.sql
var postAnalyticsQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getAnalyticsStmt, err := db.Prepare(getAnalyticsQuery)
	if err != nil {
		return nil, err
	}
	postAnalyticsStmt, err := db.Prepare(postAnalyticsQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:            db,
		getAnalytics:  getAnalyticsStmt,
		postAnalytics: postAnalyticsStmt,
	}, nil
}

type repository struct {
	db            *sql.DB
	getAnalytics  *sql.Stmt
	postAnalytics *sql.Stmt
}
