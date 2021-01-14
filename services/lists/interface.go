package lists

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
)

type Service interface {
	Create(projectID int64, name, status string, position int16) (list models.List, extErr exterrors.ExtError)
}
