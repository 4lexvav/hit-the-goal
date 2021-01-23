package tasks

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Update(taskID int, name, description string, position int16, listID int64) (task models.Task, extErr exterrors.ExtError) {
	task, extErr = srv.GetByID(taskID)
	if extErr != nil {
		return
	}

	task.Name = name
	task.Description.String = description
	task.Position = position
	task.ListID = listID

	task, err := srv.tasksDao.Update(task)
	if err != nil {
		return models.Task{}, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at updating task."))
	}

	return
}
