package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host string
	Port string
}

func GetEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	return &Env{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}
}
