package task

type Task struct {
	ID        int
	Name      string
	Completed bool
}

func NewTask(id int, name string) *Task {
	return &Task{
		ID:        id,
		Name:      name,
		Completed: false,
	}
}
