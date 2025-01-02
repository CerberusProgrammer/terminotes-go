package main

import (
	"flag"
	"log"
	"terminotes/src/commands"
	jsonmanager "terminotes/src/data/json-manager"
)

func main() {
	err := jsonmanager.CreateTable()
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	flag.Parse()
	commands.HandleCommands()
}
