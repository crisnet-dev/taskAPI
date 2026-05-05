package routes

import (
	"github.com/crisnet-dev/task-api/internal/handlers"
	"github.com/crisnet-dev/task-api/internal/middlewares"
	"github.com/gorilla/mux"
)

func RegisterTaskRoutes(mux *mux.Router) {

	handlers := handlers.NewTaskHandler()

	mux.HandleFunc("/user/task", middlewares.VerifyToken(handlers.GetAllTasksHandler)).Methods("GET")
	mux.HandleFunc("/user/task", middlewares.VerifyToken(handlers.AddTaskHandler)).Methods("POST")
	mux.HandleFunc("/user/task/{id}", middlewares.VerifyToken(handlers.DeleteTaskHandler)).Methods("DELETE")
	mux.HandleFunc("/user/task", middlewares.VerifyToken(handlers.DeleteAllTasksHandler)).Methods("DELETE")
	mux.HandleFunc("/user/task/{id}", middlewares.VerifyToken(handlers.UpdateTaskHandler)).Methods("PUT")

}
