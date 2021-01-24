package lists

import (
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) GetMaxPosition(projectID int64) (position int16, extErr exterrors.ExtError) {
	position, err := srv.listsDao.GetMaxPosition(projectID)
	if err != nil {
		extErr = exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at retrieving max list position."))
		return
	}

	return
}
