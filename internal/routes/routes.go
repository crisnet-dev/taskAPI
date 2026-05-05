package routes

import (
	"net/http"

	"github.com/crisnet-dev/task-api/internal/middlewares"
	"github.com/gorilla/mux"
)

func SetUpRoutes() *mux.Router {
	mux := mux.NewRouter()

	r := mux.PathPrefix("/api").Subrouter()

	r.Use(middlewares.SetCORS)
	r.Methods(http.MethodOptions)

	RegisterAuthRoutes(r)
	RegisterUserRoutes(r)
	RegisterTaskRoutes(r)

	return mux
}
