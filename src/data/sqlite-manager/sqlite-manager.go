package sqlitemanager

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"terminotes/src/config"
)

func checkSQLite3Installed() error {
	cmd := exec.Command("sqlite3", "--version")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("sqlite3 is not installed")
	}
	return nil
}

func CreateTable() error {
	if err := checkSQLite3Installed(); err != nil {
		return err
	}

	cmd := exec.Command("sqlite3", config.DBPath, `
    CREATE TABLE IF NOT EXISTS notes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        content TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error creating table:", err, string(output))
		return err
	}

	if string(output) == "" {
		log.Println("Table already exists")
	} else {
		log.Println("Table created successfully")
	}
	return nil
}

func CreateNote(
	title string,
	content string,
) error {
	if err := checkSQLite3Installed(); err != nil {
		return err
	}

	query := fmt.Sprintf(`
    INSERT INTO notes (title, content) VALUES ('%s', '%s');`,
		title,
		content,
	)

	cmd := exec.Command("sqlite3", config.DBPath, query)

	err := cmd.Run()
	if err != nil {
		log.Println("Error creating note:", err)
		return err
	}

	log.Println("Note created successfully")
	return nil
}

func SelectNotes(columns []string) (string, error) {
	if err := checkSQLite3Installed(); err != nil {
		return "", err
	}

	query := fmt.Sprintf("SELECT %s FROM notes;", strings.Join(columns, ", "))

	cmd := exec.Command("sqlite3", config.DBPath, query)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error selecting notes:", err, string(output))
		return "", err
	}

	return string(output), nil
}

func SelectNoteByID(id int, columns []string) (string, error) {
	if err := checkSQLite3Installed(); err != nil {
		return "", err
	}

	query := fmt.Sprintf("SELECT %s FROM notes WHERE id = %d;", strings.Join(columns, ", "), id)

	cmd := exec.Command("sqlite3", config.DBPath, query)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error selecting note by ID:", err, string(output))
		return "", err
	}

	return string(output), nil
}
