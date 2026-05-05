package routes

import (
	"github.com/crisnet-dev/task-api/internal/handlers"
	"github.com/crisnet-dev/task-api/internal/middlewares"
	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(mux *mux.Router) {

	handlers := handlers.NewTaskHandler()

	protected := mux.PathPrefix("/user/task").Subrouter()

	protected.Use(middlewares.VerifyToken)

	protected.HandleFunc("/", handlers.GetAllTasksHandler).Methods("GET")
	protected.HandleFunc("/", handlers.AddTaskHandler).Methods("POST")
	protected.HandleFunc("/{id}", handlers.DeleteTaskHandler).Methods("DELETE")
	protected.HandleFunc("/", handlers.DeleteAllTasksHandler).Methods("DELETE")
	protected.HandleFunc("/{id}", handlers.UpdateTaskHandler).Methods("PUT")

}
