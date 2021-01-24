package comments

import (
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
)

type DAO interface {
	WithTx(tx *postgres.DBQuery) DAO

	Get(taskID, size, page int) (lists []models.Comment, err error)

	GetByID(commentID int) (comment models.Comment, err error)

	Insert(comment models.Comment) (_ models.Comment, err error)

	Update(comment models.Comment) (_ models.Comment, err error)

	Delete(commentID int) (err error)
}
