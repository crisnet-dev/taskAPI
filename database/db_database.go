package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConfigDB() error {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=5432 dbname=taskSystem sslmode=disable")
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users (id) ON DELETE CASCADE,
			task_name TEXT NOT NULL,
			create_at DATE DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return err
	}

	DB = db

	return nil
}
