package lists

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Get(projectID, size, page int) (lists []models.List, extErr exterrors.ExtError) {
	lists, err := srv.listsDao.Get(projectID, size, page)
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at getting project columns"))
	}

	return lists, nil
}
