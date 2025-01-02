package search

import (
	"flag"
	"fmt"
	"log"
	"strings"
	jsonmanager "terminotes/src/data/json-manager"
	"terminotes/src/utils"
)

func SearchCommand(args []string) {
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	query := searchCmd.String("query", "", "Text to search for")
	showID := searchCmd.Bool("show-id", false, "Show ID")
	showTitle := searchCmd.Bool("show-title", false, "Show Title")
	showContent := searchCmd.Bool("show-content", false, "Show Content")
	showCreatedAt := searchCmd.Bool("show-createdAt", false, "Show Created At")

	searchCmd.Parse(args)

	if *query == "" {
		fmt.Println("Query text is required")
		searchCmd.PrintDefaults()
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

	results, err := jsonmanager.SearchNotes(*query, columns)
	if err != nil {
		log.Fatal(err)
	}

	if results == "" {
		fmt.Println("No results found")
		return
	}

	rows := strings.Split(results, "\n")
	utils.PrintTable(columns, rows)
}
