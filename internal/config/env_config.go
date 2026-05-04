package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host string
	Port string

	Issuer    string
	SecretKey string

	DatabaseHost      string
	DatabasePort      string
	DatabasePasword   string
	DatabaseUserName  string
	DatabaseTableName string
}

func GetEnv() *Env {
	return &Env{
		Host:              os.Getenv("HOST"),
		Port:              os.Getenv("PORT"),
		SecretKey:         os.Getenv("SECRET_KEY"),
		Issuer:            os.Getenv("ISSUER"),
		DatabaseHost:      os.Getenv("DATABASE_HOST"),
		DatabasePort:      os.Getenv("DATABASE_PORT"),
		DatabasePasword:   os.Getenv("DATABASE_PASSWORD"),
		DatabaseTableName: os.Getenv("DATABASE_TABLENAME"),
		DatabaseUserName:  os.Getenv("DATABASE_USERNAME"),
	}
}

func LoadConfig() error {
	err := godotenv.Load()
	return err
}
