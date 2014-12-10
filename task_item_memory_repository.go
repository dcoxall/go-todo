package todo

import (
	"errors"
	"sync"
	"sync/atomic"
)

type TaskItemMemoryRepository struct {
	persistedTasks []TaskItem
	lock           sync.RWMutex
	nextID         uint32
}

var (
	MissingTaskItemError = errors.New("Unable to find TaskItem")
	SaveTaskItemError    = errors.New("Unable to save TaskItem")
)

func (repo *TaskItemMemoryRepository) FindOne(id uint) (TaskItem, error) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()

	for _, taskItem := range repo.persistedTasks {
		if taskItem.ID == id {
			return taskItem, nil
		}
	}
	return TaskItem{}, MissingTaskItemError
}

func (repo *TaskItemMemoryRepository) Save(task *TaskItem) error {
	if task.ID > 0 {
		return SaveTaskItemError
	}

	repo.lock.Lock()
	defer repo.lock.Unlock()

	atomic.AddUint32(&repo.nextID, 1)
	task.ID = uint(atomic.LoadUint32(&repo.nextID))
	repo.persistedTasks = append(repo.persistedTasks, *task)
	return nil
}

func (repo *TaskItemMemoryRepository) FindUnComplete() (tasks []TaskItem) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()

	for _, taskItem := range repo.persistedTasks {
		if !taskItem.IsComplete() {
			tasks = append(tasks, taskItem)
		}
	}

	return
}

func (repo *TaskItemMemoryRepository) ClearAll() {
	repo.lock.Lock()
	defer repo.lock.Unlock()

	repo.persistedTasks = make([]TaskItem, 0)
}
