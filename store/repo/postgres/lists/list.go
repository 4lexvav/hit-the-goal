package lists

import (
	"context"

	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
)

type listsDAO struct {
	q *postgres.DBQuery
}

func NewListDao() DAO {
	return &listsDAO{q: postgres.GetDB().QueryContext(context.Background())}
}

func (dao listsDAO) Insert(list models.List) (_ models.List, err error) {
	_, err = dao.q.Model(&list).
		Returning("*").
		Insert()

	return list, err
}
