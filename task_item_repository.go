package todo

type TaskItemRepository interface {
	Save(*TaskItem) error
	FindOne(uint) (TaskItem, error)
	FindUnComplete() []TaskItem
}
