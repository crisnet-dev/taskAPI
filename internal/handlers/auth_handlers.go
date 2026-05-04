package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/crisnet-dev/task-api/internal/models"
	"github.com/crisnet-dev/task-api/internal/services"
	"github.com/crisnet-dev/task-api/internal/utils"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (handler *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
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
		http.StatusCreated,
	)
}

func (handler *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err)
		utils.HttpError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	accessToken, refreshToken, err := services.LoginService(user)
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
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
		http.StatusOK,
	)
}

func (handler *AuthHandler) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var refreshToken models.RefreshToken

	if err := json.NewDecoder(r.Body).Decode(&refreshToken); err != nil {
		utils.HttpError(w, "Invalid JSON", 400)
		return
	}
	defer r.Body.Close()

	accessTokenString, refreshTokenString, err := services.RefreshTokenService(refreshToken.RefreshToken)
	if err != nil {
		utils.HttpError(w, err.Error(), 401)
		return
	}

	utils.HttpResponse(
		w,
		map[string]any{
			"access_token":  accessTokenString,
			"refresh_token": refreshTokenString,
		},
		http.StatusOK,
	)

}
