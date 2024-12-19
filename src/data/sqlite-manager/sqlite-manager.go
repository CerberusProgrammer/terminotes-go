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
) (int, error) {
	if err := checkSQLite3Installed(); err != nil {
		return 0, err
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
		return 0, err
	}

	query = "SELECT * FROM notes ORDER BY id DESC LIMIT 1;"
	cmd = exec.Command("sqlite3", config.DBPath, query)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error getting last insert ID:", err, string(output))
		return 0, err
	}

	var id int
	fmt.Sscanf(strings.TrimSpace(string(output)), "%d", &id)
	return id, nil
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

func DeleteNoteByID(id int) error {
	if err := checkSQLite3Installed(); err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM notes WHERE id = %d;", id)

	cmd := exec.Command("sqlite3", config.DBPath, query)

	err := cmd.Run()
	if err != nil {
		log.Println("Error deleting note by ID:", err)
		return err
	}

	return nil
}

func DeleteAllNotes() error {
	if err := checkSQLite3Installed(); err != nil {
		return err
	}

	query := "DELETE FROM notes;"

	cmd := exec.Command("sqlite3", config.DBPath, query)

	err := cmd.Run()
	if err != nil {
		log.Println("Error deleting all notes:", err)
		return err
	}

	return nil
}

func SearchNotes(query string, columns []string) (string, error) {
	if err := checkSQLite3Installed(); err != nil {
		return "", err
	}

	query = fmt.Sprintf("SELECT %s FROM notes WHERE LOWER(title) LIKE '%%%s%%' OR LOWER(content) LIKE '%%%s%%';",
		strings.Join(columns, ", "),
		strings.ToLower(query),
		strings.ToLower(query),
	)

	cmd := exec.Command("sqlite3", config.DBPath, query)

	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("Error searching notes:", err, string(output))
		return "", err
	}

	return string(output), nil
}
