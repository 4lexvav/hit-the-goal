package projects

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Create(name string, description string) (project models.Project, extErr exterrors.ExtError) {
	project = models.Project{Name: name, Description: description}
	project, err := srv.projectsDao.Insert(project)
	if err != nil {
		return models.Project{}, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at project creation"))
	}

	return project, nil
}
