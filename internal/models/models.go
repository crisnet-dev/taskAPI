package models

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	TotalTasks int    `json:"total_tasks"`
}

type Task struct {
	ID       int    `json:"id"`
	TaskName string `json:"task_name"`
	CreateAt string `json:"create_at"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}
