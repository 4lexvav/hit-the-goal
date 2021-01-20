package lists

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) GetByID(listID int) (list models.List, extErr exterrors.ExtError) {
	list, err := srv.listsDao.GetByID(listID)
	if err != nil {
		extErr = exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at retrieving list."))
	}

	return
}
