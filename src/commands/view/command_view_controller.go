package view

import (
	"flag"
	"fmt"
	"log"
	sqlitemanager "terminotes/src/data/sqlite-manager"
	"terminotes/src/utils"
)

func ViewCommand(args []string) {
	viewCmd := flag.NewFlagSet("view", flag.ExitOnError)
	id := viewCmd.Int("id", 0, "ID of the note to view")
	showID := viewCmd.Bool("show-id", false, "Show ID")
	showTitle := viewCmd.Bool("show-title", false, "Show Title")
	showContent := viewCmd.Bool("show-content", false, "Show Content")
	showCreatedAt := viewCmd.Bool("show-createdAt", false, "Show Created At")

	viewCmd.Parse(args)

	if *id == 0 {
		fmt.Println("ID is required")
		viewCmd.PrintDefaults()
		return
	}

	columns := []string{}
	if *showID {
		columns = append(columns, "id")
	}
	if *showTitle {
		columns = append(columns, "title")
	}
	if *showContent {
		columns = append(columns, "content")
	}
	if *showCreatedAt {
		columns = append(columns, "created_at")
	}

	if len(columns) == 0 {
		columns = []string{"id", "title", "content", "created_at"}
	}

	note, err := sqlitemanager.SelectNoteByID(*id, columns)
	if err != nil {
		log.Fatal(err)
	}

	if note == "" {
		fmt.Printf("Note with ID %d not found\n", *id)
		return
	}

	utils.PrintNote(note, columns)
}
