package todo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var taskItemRepository TaskItemMemoryRepository

func testTaskRepo() TaskItemRepository {
	return &taskItemRepository
}

func TestSavingTaskItem(t *testing.T) {
	taskItem := TaskItem{
		Description: "Task description",
	}
	assert.NoError(t, testTaskRepo().Save(&taskItem))
	assert.True(t, taskItem.ID > 0)
	taskItem.Description = "Updated"
	assert.NoError(t, testTaskRepo().Save(&taskItem))
}

func TestFindingTaskItem(t *testing.T) {
	taskItem := TaskItem{
		Description: "Task description",
	}
	assert.NoError(t, testTaskRepo().Save(&taskItem))
	persistedTask, err := testTaskRepo().FindOne(taskItem.ID)
	assert.NotNil(t, persistedTask)
	assert.NoError(t, err)
}

func TestFindingNonCompletedTaskItems(t *testing.T) {
	tasks := []TaskItem{
		{
			Description: "Task description",
		},
		{
			Description: "Yet another task",
			CompletedAt: time.Now(),
		},
	}

	// Clear the in memory data store
	if memoryRepo, ok := testTaskRepo().(*TaskItemMemoryRepository); ok {
		memoryRepo.ClearAll()
	}

	for _, task := range tasks {
		assert.NoError(t, testTaskRepo().Save(&task))
	}

	unCompletedTasks := testTaskRepo().FindUnComplete()
	assert.Len(t, unCompletedTasks, 1)
	assert.False(t, unCompletedTasks[0].IsComplete())
}
