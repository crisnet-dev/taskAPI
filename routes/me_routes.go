package routes

import (
	"encoding/json"
	"net/http"

	"github.com/crisnet-dev/models"
	"github.com/crisnet-dev/services"
	"github.com/crisnet-dev/utils"
)

func ProfileRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	email := r.Context().Value("email").(string)

	userFounded, totalTasks, err := services.ProfileService(email)
	if err != nil {
		utils.HttpError(w, "User not found", http.StatusNotFound)
		return
	}

	utils.HttpResponse(
		w,
		map[string]any{
			"id":         userFounded.ID,
			"email":      userFounded.Email,
			"name":       userFounded.Name,
			"total_taks": totalTasks,
		},
		http.StatusOK,
	)
}

func DeleteAccountRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	email := r.Context().Value("email").(string)

	err := services.DeleteAccountService(email)
	if err != nil {
		utils.HttpError(w, "Some error", http.StatusInternalServerError)
		return
	}

	utils.HttpResponse(
		w,
		map[string]any{
			"message": "Account deleted",
		},
		http.StatusOK,
	)
}

func AddTaskRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		utils.HttpError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user_id := r.Context().Value("user_id").(float64)

	err := services.AddTaskService(task, user_id)
	if err != nil {
		if err.Error() == "M_C" {
			utils.HttpError(w, "Missing credentials", http.StatusBadRequest)
			return
		}

		utils.HttpError(w, "Error to add task", http.StatusInternalServerError)
		return
	}

	utils.HttpResponse(
		w,
		map[string]any{
			"message": "Task created",
		},
		http.StatusOK,
	)
}

func GetAllTasksRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user_id := r.Context().Value("user_id").(float64)

	tasks, err := services.GetAllTaskService(user_id)
	if err != nil {
		utils.HttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.HttpResponse(
		w,
		map[string]any{
			"tasks": tasks,
		},
		http.StatusOK,
	)
}

func DeleteTaskRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	task_id := r.PathValue("id")
	user_id := r.Context().Value("user_id").(float64)

	err := services.DeleteTaskService(task_id, user_id)
	if err != nil {
		utils.HttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.HttpResponse(
		w,
		map[string]any{
			"message": "Task deleted",
		},
		http.StatusOK,
	)
}

func DeleteAllTasksRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user_id := r.Context().Value("user_id").(float64)

	err := services.DeleteAllTasksService(user_id)
	if err != nil {
		utils.HttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.HttpResponse(
		w,
		map[string]any{
			"message": "Task deleted",
		},
		http.StatusOK,
	)
}

func UpdateTaskRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		utils.HttpError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	task_id := r.PathValue("id")
	user_id := r.Context().Value("user_id").(float64)

	err := services.UpdateTaskService(task.TaskName, task_id, user_id)
	if err != nil {
		utils.HttpError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.HttpResponse(
		w,
		map[string]any{
			"message": "Task updated",
		},
		http.StatusOK,
	)
}
