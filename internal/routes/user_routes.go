package routes

import (
	"github.com/crisnet-dev/task-api/internal/handlers"
	"github.com/crisnet-dev/task-api/internal/middlewares"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(mux *mux.Router) {

	handlers := handlers.NewUserHandler()

	protected := mux.PathPrefix("/user").Subrouter()

	protected.Use(middlewares.VerifyToken)

	protected.HandleFunc("/", handlers.ProfileHandler).Methods("GET")
	protected.HandleFunc("/delete", handlers.DeleteAccountHandler).Methods("DELETE")

}
