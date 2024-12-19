package commands

import (
	"terminotes/src/commands/add"
	"terminotes/src/commands/delete"
	"terminotes/src/commands/help"
	"terminotes/src/commands/list"
	"terminotes/src/commands/view"
	"terminotes/src/models"
)

var commands = []models.Command{
	{
		Name:        "add",
		Description: "Add a new note",
		Execute:     add.AddCommand,
		Error:       "Content is required",
	},
	{
		Name:        "list",
		Description: "List all notes",
		Execute:     list.ListCommand,
		Error:       "Expected 'list' command",
	},
	{
		Name:        "view",
		Description: "View a note by ID",
		Execute:     view.ViewCommand,
		Error:       "Expected 'view' command",
	},
	{
		Name:        "delete",
		Description: "Delete a note by ID or all notes",
		Execute:     delete.DeleteCommand,
		Error:       "Expected 'delete' command",
	},
	{
		Name:        "help",
		Description: "Show available commands",
		Execute:     help.HelpCommand,
		Error:       "Expected 'help' command",
	},
}

func GetCommands() []models.Command {
	return commands
}
