package comments

import (
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
	exterrors "github.com/eugeneradionov/ext-errors"
)

type Service interface {
	WithTx(tx *postgres.DBQuery) Service

	Get(taskID, size, page int) (comments []models.Comment, extErr exterrors.ExtError)

	GetByID(commentID int) (comment models.Comment, extErr exterrors.ExtError)

	Create(taskID int64, text string) (comment models.Comment, extErr exterrors.ExtError)

	Update(commentID int, text string) (comment models.Comment, extErr exterrors.ExtError)

	Delete(commentID int) (extErr exterrors.ExtError)
}
