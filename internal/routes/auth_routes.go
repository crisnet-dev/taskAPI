package routes

import (
	"github.com/crisnet-dev/task-api/internal/handlers"
	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(mux *mux.Router) {

	authHandlers := handlers.NewAuthHandler()

	mux.HandleFunc("/auth/login", authHandlers.LoginHandler).Methods("POST")
	mux.HandleFunc("/auth/signup", authHandlers.RegisterHandler).Methods("POST")
	mux.HandleFunc("/auth/refresh", authHandlers.RefreshTokenHandler).Methods("POST")

}
