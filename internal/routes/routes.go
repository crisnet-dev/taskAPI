package routes

import "net/http"

func SetUpRoutes() http.Handler {
	mux := http.NewServeMux()

	RegisterAuthRoutes(mux)
	RegisterUserRoutes(mux)
	RegisterTaskRoutes(mux)

	return mux
}
