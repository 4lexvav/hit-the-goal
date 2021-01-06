package projects

import (
	"context"

	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
)

type projectsDAO struct {
	q *postgres.DBQuery
}

func NewProjectsDao() DAO {
	return &projectsDAO{q: postgres.GetDB().QueryContext(context.Background())}
}

func (dao projectsDAO) WithTx(tx *postgres.DBQuery) DAO {
	return &projectsDAO{q: tx}
}

func (dao projectsDAO) Get(size, page int) (projects []models.Project, err error) {
	err = dao.q.Model(&projects).
		Limit(size).
		Offset(page).
		Order("created_at ASC").
		Select()

	return projects, err
}

func (dao projectsDAO) Insert(project models.Project) (_ models.Project, err error) {
	_, err = dao.q.Model(&project).
		Returning("*").
		Insert()

	return project, err
}

func (dao projectsDAO) Update(project models.Project) (_ models.Project, err error) {
	_, err = dao.q.Model(&project).
		OnConflict("(project) DO UPDATE").
		Returning("*").
		Insert()

	return project, err
}
