package comments

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Update(commentID int, text string) (comment models.Comment, extErr exterrors.ExtError) {
	comment, extErr = srv.GetByID(commentID)
	if extErr != nil {
		return
	}

	comment.Text = text

	comment, err := srv.commentsDao.Update(comment)
	if err != nil {
		return models.Comment{}, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at updating comment."))
	}

	return
}
