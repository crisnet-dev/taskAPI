package routes

import (
	"net/http"

	"github.com/crisnet-dev/task-api/internal/handlers"
	"github.com/crisnet-dev/task-api/internal/middlewares"
)

func RegisterUserRoutes(mux *http.ServeMux) {

	handlers := handlers.NewUserHandler()

	mux.HandleFunc("GET /api/user", middlewares.VerifyToken(handlers.ProfileHandler))
	mux.HandleFunc("DELETE /api/user/delete", middlewares.VerifyToken(handlers.DeleteAccountHandler))

	mux.HandleFunc("GET /api/user/task", middlewares.VerifyToken(handlers.GetAllTasksHandler))
	mux.HandleFunc("POST /api/user/task", middlewares.VerifyToken(handlers.AddTaskHandler))
	mux.HandleFunc("DELETE /api/user/task/{id}", middlewares.VerifyToken(handlers.DeleteTaskHandler))
	mux.HandleFunc("DELETE /api/user/task", middlewares.VerifyToken(handlers.DeleteAllTasksHandler))
	mux.HandleFunc("PUT /api/user/task/{id}", middlewares.VerifyToken(handlers.UpdateTaskHandler))

}
