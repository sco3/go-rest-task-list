package tasks

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: []Task{
			{ID: 0, Title: "Init"},
		},
	}
}
