package projects

import (
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Delete(id int) (extErr exterrors.ExtError) {
	err := srv.projectsDao.Delete(id)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at project deletion"))
	}

	return nil
}
