package projects

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
)

func (srv service) GetByID(id int) (project models.Project, extErr exterrors.ExtError) {
	project, err := srv.projectsDao.GetByID(id)
	if err != nil {
		return models.Project{}, exterrors.NewNotFoundError(err, "Failed at getting project")
	}

	return project, nil
}
