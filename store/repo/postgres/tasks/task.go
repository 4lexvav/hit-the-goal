package tasks

import (
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
)

const (
	stmtBaseQueryTask = "SELECT id, name, description, position, list_id, updated_at, created_at FROM tasks"
	stmtQueryTasks    = stmtBaseQueryTask + " WHERE list_id = $1 ORDER BY position ASC, created_at ASC LIMIT $2 OFFSET $3"
	stmtQueryTask     = stmtBaseQueryTask + " WHERE id = $1"
	stmtInsertTask    = "INSERT INTO tasks(name, description, position, list_id) VALUES($1, $2, $3, $4) RETURNING id, updated_at, created_at"
	stmtUpdateTask    = "UPDATE tasks SET name = $1, description = $2, position = $3, list_id = $4 WHERE id = $5 RETURNING updated_at"
	stmtDeleteTask    = "DELETE FROM tasks WHERE id = $1"
)

type tasksDAO struct {
	db *postgres.DBQuery
}

func NewTaskDao() DAO {
	return &tasksDAO{db: postgres.GetDB().NewQuery()}
}

func (dao tasksDAO) WithTx(tx *postgres.DBQuery) DAO {
	return &tasksDAO{db: tx}
}

func (dao tasksDAO) Get(listID, size, page int) (tasks []models.Task, err error) {
	rows, err := dao.db.Query(stmtQueryTasks, listID, size, size*(page-1))
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var task models.Task
		if err = rows.Scan(
			&task.ID,
			&task.Name,
			&task.Description,
			&task.Position,
			&task.ListID,
			&task.UpdatedAt,
			&task.CreatedAt); err != nil {
			return
		}

		tasks = append(tasks, task)
	}

	err = rows.Err()
	return
}

func (dao tasksDAO) GetByID(taskID int) (task models.Task, err error) {
	err = dao.db.QueryRow(stmtQueryTask, taskID).
		Scan(&task.ID, &task.Name, &task.Description, &task.Position, &task.ListID, &task.UpdatedAt, &task.CreatedAt)
	return
}

func (dao tasksDAO) Insert(task models.Task) (_ models.Task, err error) {
	if err = dao.db.QueryRow(stmtInsertTask, task.Name, task.Description, task.Position, task.ListID).
		Scan(&task.ID, &task.UpdatedAt, &task.CreatedAt); err != nil {
		return models.Task{}, err
	}

	return task, err
}

func (dao tasksDAO) Update(task models.Task) (_ models.Task, err error) {
	if err = dao.db.QueryRow(stmtUpdateTask, task.Name, task.Description, task.Position, task.ListID, task.ID).
		Scan(&task.UpdatedAt); err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (dao tasksDAO) Delete(id int) (err error) {
	stmt, err := dao.db.Prepare(stmtDeleteTask)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
