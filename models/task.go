package models

import (
	"database/sql"
	"time"

	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
)

type Task struct {
	ID          int64                   `json:"id"`
	Name        string                  `json:"name"`
	Position    int16                   `json:"position"`
	Description postgres.JsonNullString `json:"description"`

	ListID int64 `json:"list_id"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTask(name, description string, position int16, listID int64) (_ Task) {
	return Task{
		Name:     name,
		Position: position,
		Description: postgres.JsonNullString{
			NullString: sql.NullString{
				String: description,
				Valid:  true,
			},
		},
		ListID: listID,
	}
}
