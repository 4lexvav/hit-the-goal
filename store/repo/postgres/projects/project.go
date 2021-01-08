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

func (dao projectsDAO) GetByID(id int) (project models.Project, err error) {
	err = dao.q.Model(&project).
		Where("id = ?", id).
		Returning("*").
		Select()

	return project, err
}

func (dao projectsDAO) Insert(project models.Project) (_ models.Project, err error) {
	_, err = dao.q.Model(&project).
		Returning("*").
		Insert()

	return project, err
}

func (dao projectsDAO) Upsert(project models.Project) (_ models.Project, err error) {
	_, err = dao.q.Model(&project).
		OnConflict("(id) DO UPDATE").
		Returning("*").
		Insert()

	return project, err
}

func (dao projectsDAO) Delete(id int) (err error) {
	_, err = dao.q.Model((*models.Project)(nil)).
		Where("id = ?", id).
		Delete()

	return err
}
