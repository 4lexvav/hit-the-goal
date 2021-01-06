package models

import "time"

type Project struct {
	ID          uint64 `json:"id"`
	Name        string `jsong:"name"`
	Description string `jsong:"description"`

	UpdatedAt time.Time `jsong:"updated_at"`
	CreatedAt time.Time `jsong:"created_at"`
}
