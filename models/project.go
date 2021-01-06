package models

import "time"

type Project struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
