package projects

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
)

type Service interface {
	GetByID(id int) (project models.Project, extErr exterrors.ExtError)
	Get(size, page int) (projects []models.Project, extErr exterrors.ExtError)
	Create(name, description string) (project models.Project, extErr exterrors.ExtError)
	Update(id int, name, description string) (project models.Project, extErr exterrors.ExtError)
	Delete(id int) (extErr exterrors.ExtError)
}
