package tasks

import (
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
)

type DAO interface {
	WithTx(tx *postgres.DBQuery) DAO

	Get(listID, size, page int) (lists []models.Task, err error)

	GetByID(taskID int) (task models.Task, err error)

	Insert(task models.Task) (_ models.Task, err error)

	Update(task models.Task) (_ models.Task, err error)

	Delete(taskID int) (err error)
}
