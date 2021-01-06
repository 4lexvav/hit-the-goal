package projects

import (
	"github.com/4lexvav/hit-the-goal/models"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

func (srv service) Get(size, page int) (projects []models.Project, extErr exterrors.ExtError) {
	projects, err := srv.projectsDao.Get(size, page)
	if err != nil {
		return nil, exterrors.NewInternalServerErrorError(errors.Wrap(err, "get projects"))
	}

	return projects, nil
}
