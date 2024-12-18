package commands

import (
	"terminotes/src/commands/list"
	"terminotes/src/models"
)

var commands = []models.Command{
	{
		Name:        "add",
		Description: "Add a new note",
		Execute:     AddCommand,
		Error:       ErrContentRequired,
	},
	{
		Name:        "list",
		Description: "List all notes",
		Execute:     list.ListCommand,
		Error:       ErrExpectedListCommand,
	},
}
