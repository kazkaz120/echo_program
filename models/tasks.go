package models

// Task is a struct containing Task data
type Task struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Name2 string `json:"name2"`
	Name3 string `json:"name3"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTasks
func GetTasks() (tc TaskCollection) {
	tc = TaskCollection{
		[]Task{
			{1, "洗濯", "消費", "1時間"},
			{2, "勉強", "投資", "2時間"},
			{3, "ゲーム", "浪費", "2時間"},
			{4, "勉強", "投資", "3時間"},
			{5, "読書", "投資", "1時間"},
			{6, "食事", "消費", "1時間"},
		},
	}

	return
}

func CreateTasks(name string, name2 string, name3 string) (tc TaskCollection) {
	tc = TaskCollection{
		[]Task{
			{1, "洗濯", "消費", "1時間"},
			{2, "勉強", "投資", "2時間"},
			{3, "ゲーム", "浪費", "2時間"},
			{4, "勉強", "投資", "3時間"},
			{5, "読書", "投資", "1時間"},
			{6, "食事", "消費", "1時間"},
			{7, name, name2, name3},
		},
	}

	return
}
