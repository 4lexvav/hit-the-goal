package lists

import (
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
)

type DAO interface {
	WithTx(tx *postgres.DBQuery) DAO

	Get(projectID, size, page int) (lists []models.List, err error)

	GetByID(listID int) (list models.List, err error)

	GetListCount(projectID int) (count int, err error)

	GetMaxPosition(projectID int64) (position int16, err error)

	Insert(list models.List) (_ models.List, err error)

	Update(list models.List) (_ models.List, err error)

	Delete(listID int) (err error)
}
