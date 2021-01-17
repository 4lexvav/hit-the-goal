package projects

import (
	"database/sql"

	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
)

type projectsDAO struct {
	db *sql.DB
}

func NewProjectsDao() DAO {
	return &projectsDAO{db: postgres.GetDB()}
}

func (dao projectsDAO) Get(size, page int) (projects []models.Project, err error) {
	stmt := "SELECT id, name, description, updated_at, created_at FROM projects ORDER BY created_at ASC LIMIT $1 OFFSET $2"
	rows, err := dao.db.Query(stmt, size, size*(page-1))
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
	stmt := "SELECT id, name, description, updated_at, created_at FROM projects WHERE id = $1"
	if err = dao.db.QueryRow(stmt, id).Scan(&project.ID, &project.Name, &project.Description, &project.UpdatedAt, &project.CreatedAt); err != nil {
		return models.Project{}, err
	}

	return project, nil
}

func (dao projectsDAO) Insert(project models.Project) (_ models.Project, err error) {
	stmt := "INSERT INTO projects(name, description) VALUES($1, $2) RETURNING id, updated_at, created_at"
	if err = dao.db.QueryRow(stmt, project.Name, project.Description).Scan(&project.ID, &project.UpdatedAt, &project.CreatedAt); err != nil {
		return models.Project{}, err
	}

	/* stmt, err := dao.db.Prepare("INSERT INTO projects(name, description) VALUES($1, $2)")
	if err != nil {
		return models.Project{}, err
	}

	res, err := stmt.Exec(project.Name, project.Description.String)
	if err != nil {
		return models.Project{}, err
	}

	project.ID, err = res.LastInsertId()
	if err != nil {
		return project, err
	} */

	return project, nil
}

func (dao projectsDAO) Update(project models.Project) (_ models.Project, err error) {
	stmt := "UPDATE projects SET name = $1, description = $2 WHERE id = $3 RETURNING updated_at, created_at"
	if err = dao.db.QueryRow(stmt, project.Name, project.Description, project.ID).Scan(&project.UpdatedAt, &project.CreatedAt); err != nil {
		return models.Project{}, err
	}

	return project, nil
	/* stmt, err := dao.db.Prepare("UPDATE projects SET name = $1, description = $2 WHERE id = $3")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(project.Name, project.Description.String, project.ID)
	if err != nil {
		return err
	}

	return nil */
}

func (dao projectsDAO) Delete(id int) (err error) {
	stmt, err := dao.db.Prepare("DELETE FROM projects WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
