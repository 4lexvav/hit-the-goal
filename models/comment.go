package models

import (
	"time"
)

type Comment struct {
	ID   int64  `json:"id"`
	Text string `json:"text"`

	TaskID int64 `json:"task_id"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func NewComment(text string, taskID int64) (_ Comment) {
	return Comment{
		Text:   text,
		TaskID: taskID,
	}
}
