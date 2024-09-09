package staff

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/post_staff.sql
var postStaffQuery string

//go:embed sql/get_staff.sql
var getStaffQuery string

//go:embed sql/get_staff_random_by_analytics.sql
var getStaffRandomByAnalyticsQuery string

func NewRepository(db *sql.DB) (*repository, error) {

	postStaffStmt, err := db.Prepare(postStaffQuery)
	if err != nil {
		return nil, err
	}

	getStaffStmt, err := db.Prepare(getStaffQuery)
	if err != nil {
		return nil, err
	}

	getStaffRandomByAnalyticsStmt, err := db.Prepare(getStaffRandomByAnalyticsQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:                        db,
		postStaff:                 postStaffStmt,
		getStaff:                  getStaffStmt,
		getStaffRandomByAnalytics: getStaffRandomByAnalyticsStmt,
	}, nil
}

type repository struct {
	db                        *sql.DB
	postStaff                 *sql.Stmt
	getStaff                  *sql.Stmt
	getStaffRandomByAnalytics *sql.Stmt
}
