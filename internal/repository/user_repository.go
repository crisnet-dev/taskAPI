package repository

import (
	"errors"
	"log"
	"strings"

	"github.com/crisnet-dev/task-api/internal/database"
	"github.com/crisnet-dev/task-api/internal/models"
)

func AddUser(user models.User, hashedPassword string) error {
	_, err := database.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3);", user.Name, user.Email, hashedPassword)
	if err != nil {
		log.Println(err)

		if strings.Contains(strings.ToUpper(err.Error()), "UNIQUE") {
			return errors.New("UNIQUE_EMAIL_ERROR")
		}

		return errors.New("DATABASE_ERROR")
	}
	return nil
}

func DeleteUser(email string) error {
	_, err := database.DB.Exec("DELETE FROM users WHERE email = $1;", email)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := database.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1;", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)

		if strings.Contains(strings.ToUpper(err.Error()), "no rows") {
			return models.User{}, errors.New("USER_NOT_FOUND")
		}

		return models.User{}, errors.New("DATABASE_ERROR")
	}
	return user, nil
}

func FindUserByID(user_id int) (models.User, error) {
	var user models.User
	err := database.DB.QueryRow("SELECT id, name, email, password FROM users WHERE id = $1;", user_id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)

		if strings.Contains(strings.ToUpper(err.Error()), "no rows") {
			return models.User{}, errors.New("USER_NOT_FOUND")
		}

		return models.User{}, errors.New("DATABASE_ERROR")
	}
	return user, nil
}

func AddTask(task models.Task, user_id int) error {
	_, err := database.DB.Exec("INSERT INTO tasks (task_name, user_id) VALUES ($1, $2);", task.TaskName, user_id)
	if err != nil {
		log.Println(err)
		return errors.New("DATABASE_ERROR")
	}
	return nil
}

func DeleteAllTasks(user_id int) error {
	_, err := database.DB.Exec("DELETE FROM tasks WHERE user_id = $1;", user_id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DeleteTask(task_id int, user_id int) error {
	_, err := database.DB.Exec("DELETE FROM tasks WHERE id = $1 AND user_id = $2;", task_id, user_id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetTotalTasks(user_id int) (int, error) {
	var totalTasks int

	err := database.DB.QueryRow("SELECT COUNT(*) FROM tasks WHERE user_id = $1;", user_id).Scan(&totalTasks)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return totalTasks, nil
}

func GetAllTasks(user_id int) ([]models.Task, error) {
	rows, err := database.DB.Query("SELECT id, task_name, create_at FROM tasks WHERE user_id = $1;", user_id)
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

func UpdateTask(new_task_name string, task_id int, user_id int) error {
	_, err := database.DB.Exec("UPDATE tasks SET task_name = $1 WHERE id = $2 AND user_id = $3;", new_task_name, task_id, user_id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateRefreshToken(user_id int, refreshToken string) error {
	_, err := database.DB.Exec("UPDATE users SET refreshToken = $1 WHERE id = $2;", refreshToken, user_id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func WhereRefreshToken(user_id int) (models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	err := database.DB.QueryRow("SELECT refreshToken FROM users WHERE id = $1;", user_id).Scan(&refreshToken.RefreshToken)
	if err != nil {
		log.Println(err)

		if strings.Contains(strings.ToUpper(err.Error()), "no rows") {
			return models.RefreshToken{}, errors.New("REFRESH_TOKEN_NOT_FOUND")
		}

		return models.RefreshToken{}, errors.New("DATABASE_ERROR")
	}
	return refreshToken, nil
}

// func TotalTasks(email string) (int, error) {
// 	result, err := database.DB.Exec("SELECT COUNT(*) FROM tasks")
// 	if err != nil {
// 		log.Println(err)
// 		return 0, err
// 	}
// }
