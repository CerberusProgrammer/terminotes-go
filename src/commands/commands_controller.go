package commands

import (
	"flag"
	"fmt"
	"log"
	sqlitemanager "terminotes/src/data/sqlite-manager"
)

func AddCommand(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	title := addCmd.String("t", "Untitled", "Title of the note (optional)")
	content := addCmd.String("c", "", "Content of the note (required)")

	addCmd.Parse(args)

	if *content == "" {
		fmt.Println(ErrContentRequired)
		addCmd.PrintDefaults()
		return
	}

	err := sqlitemanager.CreateNote(*title, *content)
	if err != nil {
		log.Fatal(err)
	}
}

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

func HandleCommands() {
	if len(flag.Args()) < 1 {
		fmt.Println(ErrEmptyCommands)
		return
	}

	commandName := flag.Args()[0]
	args := flag.Args()[1:]

	for _, cmd := range commands {
		if cmd.Name == commandName {
			cmd.Execute(args)
			return
		}
	}

	for _, cmd := range commands {
		if cmd.Name == commandName {
			fmt.Println(cmd.Error)
			return
		}
	}

	fmt.Println(ErrUnknownCommand)
}
