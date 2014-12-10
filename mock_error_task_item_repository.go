package todo

type MockErrorTaskItemRepository struct {
	TaskItemRepository TaskItemRepository
	FindOneError       error
	SaveError          error
}

func (repo *MockErrorTaskItemRepository) FindOne(id uint) (TaskItem, error) {
	if repo.FindOneError != nil {
		return TaskItem{}, repo.FindOneError
	}
	return repo.TaskItemRepository.FindOne(id)
}

func (repo *MockErrorTaskItemRepository) Save(task *TaskItem) error {
	if repo.SaveError != nil {
		return repo.SaveError
	}
	return repo.TaskItemRepository.Save(task)
}

func (repo *MockErrorTaskItemRepository) FindUnComplete() (tasks []TaskItem) {
	return repo.TaskItemRepository.FindUnComplete()
}
