package lists

import (
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Delete(projectID, id int) (extErr exterrors.ExtError) {
	listsCount, err := srv.listsDao.GetListCount(projectID)
	if err != nil {
		return exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at counting project lists."))
	}

	if listsCount == 1 {
		return exterrors.NewInternalServerErrorError(errors.New("Cannot remove last project' list"))
	}

	err = srv.listsDao.Delete(id)
	if err != nil {
		extErr = exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at list removal."))
	}

	return
}
