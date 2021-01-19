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

func (dao listsDAO) Get(projectID, size, page int) (lists []models.List, err error) {
	stmt := "SELECT id, name, status, position, project_id, updated_at, created_at FROM lists WHERE project_id = $1 ORDER BY position ASC, created_at ASC LIMIT $2 OFFSET $3"
	rows, err := dao.db.Query(stmt, projectID, size, size*(page-1))
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var list models.List
		if err = rows.Scan(&list.ID, &list.Name, &list.Status, &list.Position, &list.ProjectID, &list.UpdatedAt, &list.CreatedAt); err != nil {
			return
		}

		lists = append(lists, list)
	}

	err = rows.Err()
	return
}

func (dao listsDAO) GetMaxPosition(projectID int64) (position int16, err error) {
	err = dao.db.QueryRow("SELECT MAX(position) FROM lists WHERE project_id = $1", projectID).Scan(&position)
	return
}

func (dao listsDAO) Insert(list models.List) (_ models.List, err error) {
	stmt := "INSERT INTO lists(name, status, position, project_id) VALUES($1, $2, $3, $4) RETURNING id, updated_at, created_at"
	if err = dao.db.QueryRow(stmt, list.Name, list.Status, list.Position, list.ProjectID).Scan(&list.ID, &list.UpdatedAt, &list.CreatedAt); err != nil {
		return models.List{}, err
	}

	return list, err
}
