package projects

import (
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
)

type DAO interface {
	WithTx(tx *postgres.DBQuery) DAO

	Get(size, page int) (projects []models.Project, err error)

	GetByID(id int) (project models.Project, err error)

	Insert(project models.Project) (_ models.Project, err error)

	Update(project models.Project) (_ models.Project, err error)

	Delete(id int) (err error)
}
