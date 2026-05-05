package models

import "github.com/golang-jwt/jwt/v5"

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

type Claims struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	jwt.RegisteredClaims
}
