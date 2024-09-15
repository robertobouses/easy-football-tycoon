package strategy

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_strategy.sql
var getStrategyQuery string

//go:embed sql/post_strategy.sql
var postStrategyQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getStrategyStmt, err := db.Prepare(getStrategyQuery)
	if err != nil {
		return nil, err
	}
	postStrategyStmt, err := db.Prepare(postStrategyQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:           db,
		getStrategy:  getStrategyStmt,
		postStrategy: postStrategyStmt,
	}, nil
}

type repository struct {
	db           *sql.DB
	getStrategy  *sql.Stmt
	postStrategy *sql.Stmt
}
