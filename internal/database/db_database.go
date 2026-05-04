package database

import (
	"database/sql"

	"github.com/crisnet-dev/task-api/internal/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConfigDB() error {

	env := config.GetEnv()
	var url string = "host=" + env.DatabaseHost + " port=" + env.DatabasePort + " user=" + env.DatabaseUserName + " password=" + env.DatabasePort + " dbname=" + env.DatabaseTableName + " sslmode=disable"

	// const url str= "host=localhost port=5432 user=postgres password=5432 dbname=taskSystem sslmode=disable"

	db, err := sql.Open("postgres", url)
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

	// _, err = db.Exec("ALTER TABLE users ADD COLUMN refreshToken TEXT;")
	// if err != nil {
	// 	return err
	// }

	DB = db

	return nil
}
