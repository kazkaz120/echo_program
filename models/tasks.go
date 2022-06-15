package models

// Task is a struct containing Task data
type Task struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Name2 string `json:"name2"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTasks
func GetTasks() (tc TaskCollection) {
	tc = TaskCollection{
		[]Task{
			{1, "洗濯", "消費"},
			{2, "勉強", "投資"},
			{3, "ゲーム", "浪費"},
			{4, "勉強", "投資"},
			{5, "読書", "投資"},
			{6, "食事", "消費"},
		},
	}

	return
}
