package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Note struct {
	ID        int
	Title     string
	Content   string
	CreatedAt string
	Completed bool
}

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

func GetNotes() ([]Note, error) {
	query := `
	SELECT id, title, content, created_at, completed
	FROM notes
	ORDER BY created_at DESC
	`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []Note
	for rows.Next() {
		var note Note
		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.Completed)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, rows.Err()
}

func GetNote(id int) (*Note, error) {
	query := `
	SELECT id, title, content, created_at, completed
	FROM notes
	WHERE id = ?
	`

	var note Note
	err := DB.QueryRow(query, id).Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &note, nil
}

func GetNode(title string) (*Note, bool, error) {
	query := `
	SELECT id, title, content, created_at, completed
	FROM notes
	WHERE title = ?
	ORDER BY created_at DESC
	`

	rows, err := DB.Query(query, title)
	if err != nil {
		return nil, false, err
	}
	defer rows.Close()

	var note *Note
	count := 0

	for rows.Next() {
		count++
		if count == 1 {
			note = &Note{}
			err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.Completed)
			if err != nil {
				return nil, false, err
			}
		}
	}

	if count == 0 {
		return nil, false, nil
	}

	duplicates := count > 1
	return note, duplicates, rows.Err()
}