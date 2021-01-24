package tasks

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Get(listID, size, page int) (tasks []models.Task, extErr exterrors.ExtError) {
	tasks, err := srv.tasksDao.Get(listID, size, page)
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at getting list tasks"))
	}

	return tasks, nil
}
