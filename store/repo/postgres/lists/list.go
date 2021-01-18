package lists

import (
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
)

type listsDAO struct {
	db *postgres.DBQuery
}

func NewListDao() DAO {
	return &listsDAO{db: postgres.GetDB().NewQuery()}
}

func (dao listsDAO) WithTx(tx *postgres.DBQuery) DAO {
	return &listsDAO{db: tx}
}

func (dao listsDAO) Insert(list models.List) (_ models.List, err error) {
	stmt := "INSERT INTO lists(name, status, position, project_id) VALUES($1, $2, $3, $4) RETURNING id, updated_at, created_at"
	if err = dao.db.QueryRow(stmt, list.Name, list.Status, list.Position, list.ProjectID).Scan(&list.ID, &list.UpdatedAt, &list.CreatedAt); err != nil {
		return models.List{}, err
	}

	return list, err
}
