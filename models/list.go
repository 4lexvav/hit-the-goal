package models

import "time"

const ListStatusActive string = "ACTIVE"

type List struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Position int16  `json:"position"`

	ProjectID int64 `json:"project_id"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
