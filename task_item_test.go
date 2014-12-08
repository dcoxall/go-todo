package todo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCoreAttributes(t *testing.T) {
	now := time.Now()
	taskItem := TaskItem{
		ID:          1,
		Description: "Task description",
		CompletedAt: now,
	}
	assert.Equal(t, taskItem.ID, 1)
	assert.Equal(t, taskItem.Description, "Task description")
	assert.Equal(t, taskItem.CompletedAt, now)
}

func TestCompletionFlag(t *testing.T) {
	taskItem := TaskItem{}
	now := time.Now()
	assert.False(t, taskItem.IsComplete())

	taskItem.CompletedAt = now.Add(-(time.Duration(1) * time.Second))
	assert.True(t, taskItem.IsComplete())

	taskItem.CompletedAt = now.Add(time.Duration(1) * time.Second)
	assert.False(t, taskItem.IsComplete())
}
