package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/crisnet-dev/models"
	"github.com/crisnet-dev/services"
	"github.com/crisnet-dev/utils"
)

func SignupRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		utils.HttpError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := services.SignupService(user)
	if err != nil {
		switch err.Error() {
		case "M_C":
			utils.HttpError(w, "Missing credentials", http.StatusBadRequest)
		case "UNIQUE_EMAIL_ERROR":
			utils.HttpError(w, "This email already just exist", http.StatusBadRequest)
		default:
			utils.HttpError(w, "Some error", http.StatusBadRequest)
		}
		return
	}

	utils.HttpResponse(
		w,
		map[string]any{
			"message": "User created",
		},
		http.StatusOK,
	)
}

func LoginRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		utils.HttpError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	token, err := services.LoginService(user)
	if err != nil {
		switch err.Error() {
		case "M_C":
			utils.HttpError(w, "Missing credentials", http.StatusBadRequest)
		case "USER_NOT_FOUND":
			utils.HttpError(w, "User not found", http.StatusNotFound)
		case "INVALID_PASSWORD":
			utils.HttpError(w, "Invalid password", http.StatusUnauthorized)
		default:
			utils.HttpError(w, "Some error", http.StatusInternalServerError)
		}
		return
	}

	utils.HttpResponse(
		w,
		map[string]any{
			"message": "Logged",
			"token":   token,
		},
		http.StatusOK,
	)
}
