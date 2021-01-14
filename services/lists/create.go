package lists

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Create(projectID int64, name, status string, position int16) (list models.List, extErr exterrors.ExtError) {
	list = models.List{Name: name, Status: status, Position: position, ProjectID: projectID}
	list, err := srv.listsDao.Insert(list)
	if err != nil {
		return models.List{}, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at list creation"))
	}

	return list, nil
}
