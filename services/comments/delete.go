package comments

import (
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Delete(id int) (extErr exterrors.ExtError) {
	err := srv.commentsDao.Delete(id)
	if err != nil {
		extErr = exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at comment removal."))
	}

	return
}
