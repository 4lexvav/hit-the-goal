package tasks

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Create(listID int64, name, description string, position int16) (task models.Task, extErr exterrors.ExtError) {
	task = models.NewTask(name, description, position, listID)
	task, err := srv.tasksDao.Insert(task)
	if err != nil {
		return models.Task{}, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at task creation"))
	}

	return task, nil
}
