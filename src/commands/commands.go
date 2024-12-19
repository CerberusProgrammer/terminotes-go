package commands

import (
	"terminotes/src/commands/add"
	"terminotes/src/commands/delete"
	"terminotes/src/commands/export"
	"terminotes/src/commands/help"
	"terminotes/src/commands/importation"
	"terminotes/src/commands/list"
	"terminotes/src/commands/search"
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
	{
		Name:        "search",
		Description: "Search notes by title or content",
		Execute:     search.SearchCommand,
		Error:       "Expected 'search' command",
	},
	{
		Name:        "export",
		Description: "Export notes to a file",
		Execute:     export.ExportCommand,
		Error:       "Expected 'export' command",
	}, {
		Name:        "import",
		Description: "Import notes from a file",
		Execute:     importation.ImportCommand,
		Error:       "Expected 'import' command",
	},
}

func GetCommands() []models.Command {
	return commands
}
