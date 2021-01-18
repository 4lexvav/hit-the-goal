package postgres

import (
	"database/sql"
)

type Querier interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Prepare(query string) (*sql.Stmt, error)
}

type DBQuery struct {
	Querier
}

func (db *DBQuery) Rollback() error {
	switch t := db.Querier.(type) {
	case *sql.Tx:
		return t.Rollback()
	}

	return nil
}

func (db *DBQuery) Commit() error {
	switch t := db.Querier.(type) {
	case *sql.Tx:
		return t.Commit()
	}

	return nil
}

func (p Postgres) NewQueryWithTx() (*DBQuery, error) {
	tx, err := p.db.Begin()
	return &DBQuery{Querier: tx}, err
}

func (p Postgres) NewQuery() *DBQuery {
	return &DBQuery{Querier: p.db}
}
