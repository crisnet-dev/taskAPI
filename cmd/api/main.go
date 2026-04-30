package main

import (
	"log"
	"net/http"

	"github.com/crisnet-dev/task-api/internal/config"
	"github.com/crisnet-dev/task-api/internal/database"
	"github.com/crisnet-dev/task-api/internal/routes"
)

func main() {
	err := database.ConfigDB()
	if err != nil {
		log.Fatal(err)
	}

	r := routes.SetUpRoutes()

	env := config.GetEnv()

	log.Printf("The server is running in: http://%s:%s\n", env.Host, env.Port)
	if err := http.ListenAndServe(env.Host+":"+env.Port, r); err != nil {
		log.Fatal(err)
	}
}
