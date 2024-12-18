package add

import (
	"flag"
	"log"
	sqlitemanager "terminotes/src/data/sqlite-manager"
)

func AddCommand(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	title := addCmd.String("t", "Untitled", "Title of the note (optional)")
	content := addCmd.String("c", "", "Content of the note (required)")

	addCmd.Parse(args)

	if *content == "" {
		addCmd.PrintDefaults()
		return
	}

	err := sqlitemanager.CreateNote(*title, *content)
	if err != nil {
		log.Fatal(err)
	}
}
