package commands

import (
	"flag"
	"fmt"
)

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
