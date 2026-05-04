package main

import (
	"log"
	"net/http"

	"github.com/crisnet-dev/task-api/internal/config"
	"github.com/crisnet-dev/task-api/internal/database"
	"github.com/crisnet-dev/task-api/internal/routes"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Error loading .env file.")
	}
	env := config.GetEnv()

	if err := database.ConfigDB(); err != nil {
		log.Fatal(err)
	}

	routes := routes.SetUpRoutes()

	log.Printf("The server is running in: http://%s:%s\n", env.Host, env.Port)
	if err := http.ListenAndServe(env.Host+":"+env.Port, routes); err != nil {
		log.Fatal(err)
	}
}
