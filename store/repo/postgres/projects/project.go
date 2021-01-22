package projects

import (
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
)

const (
	stmtBaseQueryProject = "SELECT id, name, description, updated_at, created_at FROM projects"
	stmtQueryProjects    = stmtBaseQueryProject + " ORDER BY created_at ASC LIMIT $1 OFFSET $2"
	stmtQueryProject     = stmtBaseQueryProject + " WHERE id = $1"
	stmtInsertProject    = "INSERT INTO projects(name, description) VALUES($1, $2) RETURNING id, updated_at, created_at"
	stmtUpdateProject    = "UPDATE projects SET name = $1, description = $2 WHERE id = $3 RETURNING updated_at, created_at"
	stmtDeleteProject    = "DELETE FROM projects WHERE id = $1"
)

type projectsDAO struct {
	db *postgres.DBQuery
}

func NewProjectsDao() DAO {
	return &projectsDAO{db: postgres.GetDB().NewQuery()}
}

func (dao projectsDAO) WithTx(tx *postgres.DBQuery) DAO {
	return &projectsDAO{db: tx}
}

func (dao projectsDAO) Get(size, page int) (projects []models.Project, err error) {
	rows, err := dao.db.Query(stmtQueryProjects, size, size*(page-1))
	if err != nil {
		return projects, err
	}
	defer rows.Close()

	for rows.Next() {
		var proj models.Project
		if err = rows.Scan(&proj.ID, &proj.Name, &proj.Description, &proj.UpdatedAt, &proj.CreatedAt); err != nil {
			return projects, err
		}

		projects = append(projects, proj)
	}

	err = rows.Err()
	if err != nil {
		return projects, err
	}

	return projects, nil
}

func (dao projectsDAO) GetByID(id int) (project models.Project, err error) {
	if err = dao.db.QueryRow(stmtQueryProject, id).
		Scan(&project.ID, &project.Name, &project.Description, &project.UpdatedAt, &project.CreatedAt); err != nil {
		return
	}
	return
}

func (dao projectsDAO) Insert(project models.Project) (_ models.Project, err error) {
	if err = dao.db.QueryRow(stmtInsertProject, project.Name, project.Description).
		Scan(&project.ID, &project.UpdatedAt, &project.CreatedAt); err != nil {
		return models.Project{}, err
	}
	return project, nil
}

func (dao projectsDAO) Update(project models.Project) (_ models.Project, err error) {
	if err = dao.db.QueryRow(stmtUpdateProject, project.Name, project.Description, project.ID).
		Scan(&project.UpdatedAt, &project.CreatedAt); err != nil {
		return models.Project{}, err
	}

	return project, nil
}

func (dao projectsDAO) Delete(id int) (err error) {
	stmt, err := dao.db.Prepare(stmtDeleteProject)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
