package routes

import (
	"github.com/crisnet-dev/task-api/internal/handlers"
	"github.com/crisnet-dev/task-api/internal/middlewares"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(mux *mux.Router) {

	handlers := handlers.NewUserHandler()

	mux.HandleFunc("/user", middlewares.VerifyToken(handlers.ProfileHandler)).Methods("GET")
	mux.HandleFunc("/user/delete", middlewares.VerifyToken(handlers.DeleteAccountHandler)).Methods("DELETE")

}
