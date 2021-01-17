package models

import (
	"database/sql"
	"time"

	"github.com/4lexvav/hit-the-goal/store/repo/postgres"
)

type Project struct {
	ID          int64                   `json:"id"`
	Name        string                  `json:"name"`
	Description postgres.JsonNullString `json:"description"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProject(name, description string) (_ Project) {
	return Project{
		Name: name,
		Description: postgres.JsonNullString{
			NullString: sql.NullString{
				String: description,
				Valid:  true,
			},
		},
	}
}
