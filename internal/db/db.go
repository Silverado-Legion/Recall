package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init() error {
	var err error
	DB, err = sql.Open("sqlite", "./recall.db")
	if err != nil {
		return err
	}

	return createTables()
}

func createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		completed INTEGER DEFAULT 0
	);
	`

	_, err := DB.Exec(query)
	return err
}

func AddNote(title, content string) error {
	query := `
	INSERT INTO notes (title, content)
	VALUES (?, ?)
	`

	_, err := DB.Exec(query, title, content)
	return err
}