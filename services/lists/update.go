package lists

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Update(listID int, name, status string, position int16) (list models.List, extErr exterrors.ExtError) {
	list, extErr = srv.GetByID(listID)
	if extErr != nil {
		return
	}

	list.Name = name
	list.Status = status
	list.Position = position

	list, err := srv.listsDao.Update(list)
	if err != nil {
		return models.List{}, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at updating list."))
	}

	return
}
