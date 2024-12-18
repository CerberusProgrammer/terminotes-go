package commands

import (
	"terminotes/src/commands/add"
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
}
