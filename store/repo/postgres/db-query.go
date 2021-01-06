package postgres

import (
	"context"
	"errors"

	"github.com/4lexvav/hit-the-goal/logger"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type DBModel interface {
	Model(model ...interface{}) *orm.Query
	Exec(query interface{}, params ...interface{}) (pg.Result, error)
}

type DBQuery struct {
	DBModel
	completed bool
}

func (q *DBQuery) Rollback() error {
	switch t := q.DBModel.(type) {
	case *pg.Tx:
		if !q.completed {
			return t.Rollback()
		}
		return nil
	}

	return errors.New("rollback failed: not in Tx")
}

func (q *DBQuery) RollbackTx(info string) {
	if err := q.Rollback(); err != nil {
		logger.Get().Errorw("failed to rollback transaction", "error", err, "info", info)
	}
}

func (q *DBQuery) Commit() error {
	switch t := q.DBModel.(type) { // nolint:gocritic
	case *pg.Tx:
		if !q.completed {
			q.completed = true

			return t.Commit()
		}
	}

	return nil
}

func (p Postgres) NewTXContext(ctx context.Context) (*DBQuery, error) {
	tx, err := p.conn.WithContext(ctx).Begin()
	return &DBQuery{DBModel: tx}, err
}

func (p Postgres) QueryContext(ctx context.Context) *DBQuery {
	return &DBQuery{DBModel: p.conn.WithContext(ctx)}
}
