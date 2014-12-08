package todo

import (
	"time"
)

type TaskItem struct {
	ID          uint      `json:"id,omitempty"`
	Description string    `json:"description"`
	CompletedAt time.Time `json:"completed_at,omitempty"`
}

func (task *TaskItem) IsComplete() bool {
	return !task.CompletedAt.IsZero() &&
		task.CompletedAt.Before(time.Now())
}
