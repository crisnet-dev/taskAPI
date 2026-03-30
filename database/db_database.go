package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func ConfigDB() error {
	db, err := sql.Open("sqlite", "file:database.db")
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			task_name TEXT NOT NULL,
			create_at DATE DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY(user_id) REFERENCES users (id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return err
	}

	DB = db

	return nil
}
