package tasks

import (
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
	exterrors "github.com/eugeneradionov/ext-errors"
)

type Service interface {
	WithTx(tx *postgres.DBQuery) Service

	Get(taskID, size, page int) (tasks []models.Task, extErr exterrors.ExtError)

	GetByID(taskID int) (task models.Task, extErr exterrors.ExtError)

	Create(taskID int64, name, description string, position int16) (task models.Task, extErr exterrors.ExtError)

	Update(taskID int, name, description string, position int16, listID int64) (task models.Task, extErr exterrors.ExtError)

	Delete(taskID int) (extErr exterrors.ExtError)
}
