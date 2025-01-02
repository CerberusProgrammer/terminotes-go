package jsonmanager

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"terminotes/src/config"
	"time"
)

type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

var notes []Note

func loadNotes() error {
	file, err := ioutil.ReadFile(config.DBPath)
	if err != nil {
		if os.IsNotExist(err) {
			notes = []Note{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &notes)
}

func saveNotes() error {
	file, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(config.DBPath, file, 0644)
}

func CreateTable() error {
	if _, err := os.Stat(config.DBPath); os.IsNotExist(err) {
		return saveNotes()
	}
	return nil
}

func CreateNote(
	title string,
	content string,
) (int, error) {
	if err := loadNotes(); err != nil {
		return 0, err
	}

	id := len(notes) + 1
	note := Note{
		ID:        id,
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}
	notes = append(notes, note)

	if err := saveNotes(); err != nil {
		return 0, err
	}

	return id, nil
}

func CreateNoteWithoutID(
	title string,
	content string,
) error {
	_, err := CreateNote(title, content)
	return err
}

func SelectNotes(columns []string) (string, error) {
	if err := loadNotes(); err != nil {
		return "", err
	}

	var result string
	for _, note := range notes {
		result += fmt.Sprintf("%d | %s | %s | %s\n", note.ID, note.Title, note.Content, note.CreatedAt)
	}

	return result, nil
}

func SelectNoteByID(id int, columns []string) (string, error) {
	if err := loadNotes(); err != nil {
		return "", err
	}

	for _, note := range notes {
		if note.ID == id {
			return fmt.Sprintf("%d | %s | %s | %s\n", note.ID, note.Title, note.Content, note.CreatedAt), nil
		}
	}

	return "", fmt.Errorf("note with ID %d not found", id)
}

func DeleteNoteByID(id int) error {
	if err := loadNotes(); err != nil {
		return err
	}

	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			return saveNotes()
		}
	}

	return fmt.Errorf("note with ID %d not found", id)
}

func DeleteAllNotes() error {
	notes = []Note{}
	return saveNotes()
}

func SearchNotes(searchQuery string, columns []string) (string, error) {
	if err := loadNotes(); err != nil {
		return "", err
	}

	var result string
	for _, note := range notes {
		if contains(strings.ToLower(note.Title), strings.ToLower(searchQuery)) || contains(strings.ToLower(note.Content), strings.ToLower(searchQuery)) {
			result += fmt.Sprintf("%d | %s | %s | %s\n", note.ID, note.Title, note.Content, note.CreatedAt)
		}
	}

	return result, nil
}

func SelectAllNotes() (string, error) {
	return SelectNotes(nil)
}

func ExecuteSQL(query string) error {
	return fmt.Errorf("ExecuteSQL is not supported for JSON storage")
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}
