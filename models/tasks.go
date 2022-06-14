package models

// Task is a struct containing Task data
type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTasks
func GetTasks() (tc TaskCollection) {
	tc = TaskCollection{
		[]Task{
			{1, "洗濯"},
			{2, "勉強"},
		},
	}

	return
}
