package team

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/get_rival.sql
var getRivalQuery string

//go:embed sql/post_rival.sql
var postRivalQuery string

func NewRepository(db *sql.DB) (*repository, error) {
	getRivalStmt, err := db.Prepare(getRivalQuery)
	if err != nil {
		return nil, err
	}
	postRivalStmt, err := db.Prepare(postRivalQuery)
	if err != nil {
		return nil, err
	}

	return &repository{
		db:        db,
		getRival:  getRivalStmt,
		postRival: postRivalStmt,
	}, nil
}

type repository struct {
	db        *sql.DB
	getRival  *sql.Stmt
	postRival *sql.Stmt
}
