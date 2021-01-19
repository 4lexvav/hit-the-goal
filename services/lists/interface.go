package lists

import (
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
	exterrors "github.com/eugeneradionov/ext-errors"
)

type Service interface {
	WithTx(tx *postgres.DBQuery) Service

	Get(projectID, size, page int) (lists []models.List, extErr exterrors.ExtError)

	GetMaxPosition(projectID int64) (position int16, extErr exterrors.ExtError)

	Create(projectID int64, name, status string, position int16) (list models.List, extErr exterrors.ExtError)
}
