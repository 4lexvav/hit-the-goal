package lists

import (
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
	exterrors "github.com/eugeneradionov/ext-errors"
)

type Service interface {
	WithTx(tx *postgres.DBQuery) Service

	Create(projectID int64, name, status string, position int16) (list models.List, extErr exterrors.ExtError)
}
