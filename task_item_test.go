package todo

import (
	"encoding/json"
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

func TestJSONDecoding(t *testing.T) {
	taskItemJson := `{
		"id": 1,
		"description": "Task description",
		"completed_at": "1990-01-23T05:30:00Z"
	}`
	taskItem := TaskItem{}
	err := json.Unmarshal([]byte(taskItemJson), &taskItem)
	assert.NoError(t, err)
	assert.Equal(t, taskItem.ID, 1)
	assert.Equal(t, taskItem.Description, "Task description")
	expectedTime, _ := time.Parse(
		"2 Jan 2006 @ 15:04 MST",
		"23 Jan 1990 @ 5:30 UTC",
	)
	assert.Equal(t, taskItem.CompletedAt, expectedTime)
}

func TestJSONEncoding(t *testing.T) {
	completedTime, _ := time.Parse(
		"2 Jan 2006 @ 15:04 MST",
		"23 Jan 1990 @ 5:30 UTC",
	)
	taskItem := TaskItem{
		ID:          1,
		Description: "Task description",
		CompletedAt: completedTime,
	}
	json, err := json.MarshalIndent(&taskItem, "\t", "\t")
	assert.NoError(t, err)
	expectedOutput := `{
		"id": 1,
		"description": "Task description",
		"completed_at": "1990-01-23T05:30:00Z"
	}`
	assert.Equal(t, string(json), expectedOutput)
}
