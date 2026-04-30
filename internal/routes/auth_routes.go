package routes

import (
	"net/http"

	"github.com/crisnet-dev/task-api/internal/handlers"
)

func RegisterAuthRoutes(mux *http.ServeMux) {

	authHandlers := handlers.NewAuthHandler()

	mux.HandleFunc("POST /api/auth/login", authHandlers.LoginHandler)
	mux.HandleFunc("POST /api/auth/signup", authHandlers.RegisterHandler)
	mux.HandleFunc("POST /api/auth/refresh", authHandlers.RefreshTokenHandler)

}
