package projects

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Update(id int, name string, description string) (project models.Project, extErr exterrors.ExtError) {
	project, extErr = srv.GetByID(id)
	if extErr != nil {
		return models.Project{}, extErr
	}

	project.Name = name
	project.Description.String = description

	project, err := srv.projectsDao.Update(project)
	if err != nil {
		return models.Project{}, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at updating project"))
	}

	return project, nil
}
