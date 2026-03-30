package repository

import (
	"errors"
	"log"
	"strings"

	"github.com/crisnet-dev/database"
	"github.com/crisnet-dev/models"
)

func AddUser(user models.User, hashedPassword string) error {
	_, err := database.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, hashedPassword)
	if err != nil {
		log.Println(err)

		if strings.Contains(err.Error(), "UNIQUE") {
			return errors.New("UNIQUE_EMAIL_ERROR")
		}

		return errors.New("DATABASE_ERROR")
	}
	return nil
}

func DeleteUser(email string) error {
	_, err := database.DB.Exec("DELETE FROM users WHERE email = ?", email)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := database.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)

		if strings.Contains(err.Error(), "no rows") {
			return models.User{}, errors.New("USER_NOT_FOUND")
		}

		return models.User{}, errors.New("DATABASE_ERROR")
	}
	return user, nil
}

func AddTask(task models.Task, user_id float64) error {
	_, err := database.DB.Exec("INSERT INTO tasks (task_name, user_id) VALUES (?, ?)", task.TaskName, user_id)
	if err != nil {
		log.Println(err)
		return errors.New("DATABASE_ERROR")
	}
	return nil
}

func DeleteAllTasks(user_id float64) error {
	_, err := database.DB.Exec("DELETE FROM tasks WHERE user_id = ?", user_id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteTask(task_id int, user_id float64) error {
	_, err := database.DB.Exec("DELETE FROM tasks WHERE id = ? AND user_id = ?", task_id, user_id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetAllTasks(user_id float64) ([]models.Task, error) {
	rows, err := database.DB.Query("SELECT id, task_name, create_at FROM tasks WHERE user_id = ?", user_id)
	if err != nil {
		log.Println(err)
		return []models.Task{}, err
	}

	var tasks []models.Task = []models.Task{}

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.TaskName, &task.CreateAt); err != nil {
			log.Println(err)
			return []models.Task{}, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func UpdateTask(new_task_name string, task_id int, user_id float64) error {
	_, err := database.DB.Exec("UPDATE tasks SET task_name = ? WHERE id = ? AND user_id = ?", new_task_name, task_id, user_id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// func TotalTasks(email string) (int, error) {
// 	result, err := database.DB.Exec("SELECT COUNT(*) FROM tasks")
// 	if err != nil {
// 		log.Println(err)
// 		return 0, err
// 	}
// }
