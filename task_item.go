package todo

import (
	"time"
)

type TaskItem struct {
	ID          uint
	Description string
	CompletedAt time.Time
}

func (task *TaskItem) IsComplete() bool {
	return !task.CompletedAt.IsZero() &&
		task.CompletedAt.Before(time.Now())
}
