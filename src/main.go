package main

import (
	"database/sql"
	"log"
	sqlitemanager "terminotes/src/data/sqlite-manager"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./terminotes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = sqlitemanager.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}
}
