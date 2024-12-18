package list

import (
	"flag"
	"fmt"
	"log"
	sqlitemanager "terminotes/src/data/sqlite-manager"
)

func ListCommand(args []string) {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	showID := listCmd.Bool("id", false, "Show ID")
	showTitle := listCmd.Bool("title", false, "Show Title")
	showContent := listCmd.Bool("content", false, "Show Content")
	showCreatedAt := listCmd.Bool("createdAt", false, "Show Created At")

	listCmd.Parse(args)

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

	notes, err := sqlitemanager.SelectNotes(columns)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ID | Title | Content | Created At")
	fmt.Println(notes)
}
