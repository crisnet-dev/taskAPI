package routes

import "github.com/gorilla/mux"

func SetUpRoutes() *mux.Router {
	mux := mux.NewRouter()

	r := mux.PathPrefix("/api").Subrouter()

	RegisterAuthRoutes(r)
	RegisterUserRoutes(r)
	RegisterTaskRoutes(r)

	return mux
}
