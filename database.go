package main

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func initDatabase() error {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./monopage.db"
	}

	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS pages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		slug TEXT UNIQUE NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		edit_token TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = db.Exec(createTableSQL)
	return err
}