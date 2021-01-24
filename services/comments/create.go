package comments

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Create(taskID int64, text string) (comment models.Comment, extErr exterrors.ExtError) {
	comment = models.NewComment(text, taskID)
	comment, err := srv.commentsDao.Insert(comment)
	if err != nil {
		return models.Comment{}, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at comment creation"))
	}

	return comment, nil
}
