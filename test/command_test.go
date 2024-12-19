package models

import (
	"terminotes/src/models"
	"testing"
)

func TestCommand_Execute(t *testing.T) {
	executed := false
	cmd := models.Command{
		Execute: func(args []string) {
			executed = true
		},
	}

	cmd.Execute(nil)

	if !executed {
		t.Errorf("expected command to be executed")
	}
}

func TestCommand_Error(t *testing.T) {
	cmd := models.Command{
		Name:  "test",
		Error: "test error",
	}

	if cmd.Name != "test" {
		t.Errorf("expected name to be 'test', got '%s'", cmd.Name)
	}

	if cmd.Error != "test error" {
		t.Errorf("expected error to be 'test error', got '%s'", cmd.Error)
	}
}
