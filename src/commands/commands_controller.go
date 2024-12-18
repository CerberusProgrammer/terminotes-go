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
