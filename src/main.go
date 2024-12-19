package main

import (
	"flag"
	"log"
	"terminotes/src/commands"
	sqlitemanager "terminotes/src/data/sqlite-manager"
)

func main() {
	err := sqlitemanager.CreateTable()
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	/*
		- flag.Parse(): Analiza los argumentos de la línea de comandos y los almacena internamente.
		- flag.Args(): Devuelve los argumentos no procesados como una lista de cadenas.
		- commands.HandleCommands(): Accede a flag.Args() para determinar qué comando se pasó y maneja la lógica correspondiente.
	*/
	flag.Parse()
	commands.HandleCommands()
}
