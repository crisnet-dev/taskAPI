package main

import (
	"log"
	"net/http"

	"github.com/crisnet-dev/database"
	"github.com/crisnet-dev/middlewares"
	"github.com/crisnet-dev/routes"
)

func main() {
	err := database.ConfigDB()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("POST /api/auth/login", routes.LoginRouter)
	http.HandleFunc("POST /api/auth/signup", routes.SignupRouter)

	http.HandleFunc("GET /api/me", middlewares.VerifyToken(routes.ProfileRouter))
	http.HandleFunc("DELETE /api/me/delete", middlewares.VerifyToken(routes.DeleteAccountRouter))

	http.HandleFunc("GET /api/me/task", middlewares.VerifyToken(routes.GetAllTasksRouter))
	http.HandleFunc("POST /api/me/task", middlewares.VerifyToken(routes.AddTaskRouter))
	http.HandleFunc("DELETE /api/me/task/{id}", middlewares.VerifyToken(routes.DeleteTaskRouter))
	http.HandleFunc("DELETE /api/me/task", middlewares.VerifyToken(routes.DeleteAllTasksRouter))
	http.HandleFunc("PUT /api/me/task/{id}", middlewares.VerifyToken(routes.UpdateTaskRouter))

	log.Println("The server is running...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
