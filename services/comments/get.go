package comments

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Get(taskID, size, page int) (comments []models.Comment, extErr exterrors.ExtError) {
	comments, err := srv.commentsDao.Get(taskID, size, page)
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at getting list comments"))
	}

	return comments, nil
}
