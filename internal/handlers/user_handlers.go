package handlers

import (
	"net/http"

	"github.com/crisnet-dev/task-api/internal/services"
	"github.com/crisnet-dev/task-api/internal/utils"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (handler *UserHandler) ProfileHandler(w http.ResponseWriter, r *http.Request) {
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

func (handler *UserHandler) DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
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
