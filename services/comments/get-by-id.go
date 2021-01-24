package comments

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) GetByID(commentID int) (comment models.Comment, extErr exterrors.ExtError) {
	comment, err := srv.commentsDao.GetByID(commentID)
	if err != nil {
		extErr = exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at retrieving comment"))
	}

	return
}
