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

}
