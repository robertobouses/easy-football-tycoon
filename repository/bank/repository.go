package bank

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/post_transactions.sql
var postTransactionsQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	postTransactionsStmt, err := db.Prepare(postTransactionsQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:               db,
		postTransactions: postTransactionsStmt,
	}, nil
}

type repository struct {
	db               *sql.DB
	postTransactions *sql.Stmt
}
