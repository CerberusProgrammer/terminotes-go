package sqlitemanager

import (
	"database/sql"
	"fmt"
)

func CreateTable(db *sql.DB) error {
	query := `
        CREATE TABLE IF NOT EXISTS notes (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT,
            content TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `
	_, err := db.Exec(query)

	if err != nil {
		fmt.Println("Error creating table: ", err)
	} else {
		fmt.Println("Table created successfully")
	}

	return err
}
