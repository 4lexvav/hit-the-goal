package tasks

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) GetByID(taskID int) (task models.Task, extErr exterrors.ExtError) {
	task, err := srv.tasksDao.GetByID(taskID)
	if err != nil {
		extErr = exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at retrieving task."))
	}

	return
}
