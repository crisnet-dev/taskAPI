package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/crisnet-dev/task-api/internal/models"
	"github.com/crisnet-dev/task-api/internal/services"
	"github.com/crisnet-dev/task-api/internal/utils"
	"github.com/gorilla/mux"
)

type TaskHandler struct{}

func NewTaskHandler() *TaskHandler {
	return &TaskHandler{}
}

func (handler *TaskHandler) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		utils.HttpError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user_id_string := r.Context().Value("user_id").(string)
	user_id, err := strconv.Atoi(user_id_string)
	if err != nil {
		utils.HttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.AddTaskService(task, user_id)
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
		http.StatusCreated,
	)
}

func (handler *TaskHandler) GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user_id_string := r.Context().Value("user_id").(string)
	user_id, err := strconv.Atoi(user_id_string)
	if err != nil {
		utils.HttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

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

func (handler *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	task_id := vars["id"]
	user_id_string := r.Context().Value("user_id").(string)
	user_id, err := strconv.Atoi(user_id_string)
	if err != nil {
		utils.HttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.DeleteTaskService(task_id, user_id)
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

func (handler *TaskHandler) DeleteAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user_id_string := r.Context().Value("user_id").(string)
	user_id, err := strconv.Atoi(user_id_string)
	if err != nil {
		utils.HttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.DeleteAllTasksService(user_id)
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

func (handler *TaskHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		utils.HttpError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)

	task_id := vars["id"]
	user_id_string := r.Context().Value("user_id").(string)
	user_id, err := strconv.Atoi(user_id_string)
	if err != nil {
		utils.HttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.UpdateTaskService(task.TaskName, task_id, user_id)
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
