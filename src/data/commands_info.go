package data

import "terminotes/src/models"

var CommandsInfo = []models.Command{
	{
		Name:        "add",
		Description: "Add a new note",
		Error:       "Content is required",
	},
	{
		Name:        "list",
		Description: "List all notes",
		Error:       "Expected 'list' command",
	},
	{
		Name:        "view",
		Description: "View a note by ID",
		Error:       "Expected 'view' command",
	},
	{
		Name:        "delete",
		Description: "Delete a note by ID or all notes",
		Error:       "Expected 'delete' command",
	},
	{
		Name:        "help",
		Description: "Show available commands",
		Error:       "Expected 'help' command",
	},
	{
		Name:        "search",
		Description: "Search notes by title or content",
		Error:       "Expected 'search' command",
	},
	{
		Name:        "export",
		Description: "Export notes to a file",
		Error:       "Expected 'export' command",
	},
	{
		Name:        "import",
		Description: "Import notes from a file",
		Error:       "Expected 'import' command",
	},
}
