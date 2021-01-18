package projects

import (
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/services/lists"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
	exterrors "github.com/eugeneradionov/ext-errors"
	"github.com/pkg/errors"
)

const defaultListName string = "TODO"

func (srv service) Create(name string, description string, listsSrv lists.Service) (project models.Project, extErr exterrors.ExtError) {
	tx, err := postgres.GetDB().NewQueryWithTx()
	if err != nil {
		return models.Project{}, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed initializing transaction"))
	}
	defer tx.Rollback()

	project = models.NewProject(name, description)
	project, err = srv.projectsDao.WithTx(tx).Insert(project)
	if err != nil {
		return models.Project{}, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Failed at project creation"))
	}

	if _, extErrList := listsSrv.WithTx(tx).Create(project.ID, defaultListName, models.ListStatusActive, 0); extErrList != nil {
		return models.Project{}, extErrList
	}

	err = tx.Commit()
	if err != nil {
		return models.Project{}, exterrors.NewInternalServerErrorError(errors.Wrap(err, "Error commiting transaction"))
	}

	return project, nil
}
