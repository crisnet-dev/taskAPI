package services

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/crisnet-dev/task-api/internal/models"
	"github.com/crisnet-dev/task-api/internal/repository"
)

func DeleteAccountService(email string) error {
	err := repository.DeleteUser(email)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func ProfileService(email string) (models.User, int, error) {
	userFounded, err := repository.FindUserByEmail(email)
	if err != nil {
		return models.User{}, 0, err
	}

	total_tasks, err := repository.GetTotalTasks(userFounded.ID)
	return userFounded, total_tasks, nil
}

func AddTaskService(task models.Task, user_id float64) error {
	if strings.TrimSpace(task.TaskName) == "" {
		return errors.New("M_C")
	}

	repository.AddTask(task, user_id)
	return nil
}

func GetAllTaskService(user_id float64) ([]models.Task, error) {
	tasks, err := repository.GetAllTasks(user_id)
	if err != nil {
		return []models.Task{}, err
	}
	return tasks, nil
}

func DeleteAllTasksService(user_id float64) error {
	if err := repository.DeleteAllTasks(user_id); err != nil {
		return err
	}
	return nil
}

func DeleteTaskService(task_id string, user_id float64) error {
	task_id_int, err := strconv.Atoi(task_id)
	if err != nil {
		return err
	}

	if err := repository.DeleteTask(task_id_int, user_id); err != nil {
		return err
	}
	return nil
}

func UpdateTaskService(new_task_name string, task_id string, user_id float64) error {
	task_id_int, err := strconv.Atoi(task_id)
	if err != nil {
		return err
	}

	err = repository.UpdateTask(new_task_name, task_id_int, user_id)
	if err != nil {
		return err
	}
	return nil
}
