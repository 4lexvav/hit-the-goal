package projects

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
)

type Service interface {
	Get(size, page int) (projects []models.Project, extErr exterrors.ExtError)
}
