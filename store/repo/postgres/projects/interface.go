package projects

import (
	"github.com/4lexvav/hit-the-goal/models"
)

type DAO interface {
	Get(size, page int) (projects []models.Project, err error)

	GetByID(id int) (project models.Project, err error)

	Insert(project models.Project) (_ models.Project, err error)

	Update(project models.Project) (_ models.Project, err error)

	Delete(id int) (err error)
}
