package delete

import (
	"flag"
	"fmt"
	"log"
	sqlitemanager "terminotes/src/data/sqlite-manager"
)

func DeleteCommand(args []string) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deleteCmd.Int("id", 0, "ID of the note to delete")
	all := deleteCmd.Bool("all", false, "Delete all notes")

	deleteCmd.Parse(args)

	if *all {
		err := sqlitemanager.DeleteAllNotes()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("All notes deleted successfully")
		return
	}

	if *id == 0 {
		fmt.Println("ID is required")
		deleteCmd.PrintDefaults()
		return
	}

	err := sqlitemanager.DeleteNoteByID(*id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Note %d deleted successfully\n", *id)
}
