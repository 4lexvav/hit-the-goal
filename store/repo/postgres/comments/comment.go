package comments

import (
	"github.com/4lexvav/hit-the-goal/models"
	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
)

const (
	stmtBaseQueryComment = "SELECT id, text, task_id, updated_at, created_at FROM comments"
	stmtQueryComments    = stmtBaseQueryComment + " WHERE task_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3"
	stmtQueryComment     = stmtBaseQueryComment + " WHERE id = $1"
	stmtInsertComment    = "INSERT INTO comments(text, task_id) VALUES($1, $2) RETURNING id, updated_at, created_at"
	stmtUpdateComment    = "UPDATE comments SET text = $1 WHERE id = $2 RETURNING updated_at"
	stmtDeleteComment    = "DELETE FROM comments WHERE id = $1"
)

type commentsDAO struct {
	db *postgres.DBQuery
}

func NewCommentDao() DAO {
	return &commentsDAO{db: postgres.GetDB().NewQuery()}
}

func (dao commentsDAO) WithTx(tx *postgres.DBQuery) DAO {
	return &commentsDAO{db: tx}
}

func (dao commentsDAO) Get(taskID, size, page int) (comments []models.Comment, err error) {
	rows, err := dao.db.Query(stmtQueryComments, taskID, size, size*(page-1))
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		if err = rows.Scan(
			&comment.ID,
			&comment.Text,
			&comment.TaskID,
			&comment.UpdatedAt,
			&comment.CreatedAt); err != nil {
			return
		}

		comments = append(comments, comment)
	}

	err = rows.Err()
	return
}

func (dao commentsDAO) GetByID(commentID int) (comment models.Comment, err error) {
	err = dao.db.QueryRow(stmtQueryComment, commentID).
		Scan(&comment.ID, &comment.Text, &comment.TaskID, &comment.UpdatedAt, &comment.CreatedAt)
	return
}

func (dao commentsDAO) Insert(comment models.Comment) (_ models.Comment, err error) {
	if err = dao.db.QueryRow(stmtInsertComment, comment.Text, comment.TaskID).
		Scan(&comment.ID, &comment.UpdatedAt, &comment.CreatedAt); err != nil {
		return models.Comment{}, err
	}

	return comment, err
}

func (dao commentsDAO) Update(comment models.Comment) (_ models.Comment, err error) {
	if err = dao.db.QueryRow(stmtUpdateComment, comment.Text, comment.ID).
		Scan(&comment.UpdatedAt); err != nil {
		return models.Comment{}, err
	}

	return comment, nil
}

func (dao commentsDAO) Delete(id int) (err error) {
	stmt, err := dao.db.Prepare(stmtDeleteComment)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
