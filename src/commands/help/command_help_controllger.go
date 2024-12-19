package help

import (
	"fmt"
	"terminotes/src/data"
)

func HelpCommand(args []string) {
	fmt.Println("Available commands:")

	for _, command := range data.CommandsInfo {
		fmt.Printf("%s - %s\n", command.Name, command.Description)
	}
}
