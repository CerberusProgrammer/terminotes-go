package export

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	jsonmanager "terminotes/src/data/json-manager"
)

func ExportCommand(args []string) {
	exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
	format := exportCmd.String("format", "", "Format to export (sql, json, txt)")

	exportCmd.Parse(args)

	if *format == "" {
		fmt.Println("Format is required")
		exportCmd.PrintDefaults()
		return
	}

	notes, err := jsonmanager.SelectAllNotes()
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(*format) {
	case "sql":
		exportToSQL(notes)
	case "json":
		exportToJSON(notes)
	case "txt":
		exportToTXT(notes)
	default:
		fmt.Println("Unsupported format. Supported formats are: sql, json, txt")
	}
}

func exportToSQL(notes string) {
	file, err := os.Create("notes.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(notes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Notes exported to notes.sql")
}

func exportToJSON(notes string) {
	var noteList []map[string]interface{}
	rows := strings.Split(notes, "\n")
	for _, row := range rows {
		if row == "" {
			continue
		}
		columns := strings.Split(row, "|")
		note := map[string]interface{}{
			"id":         columns[0],
			"title":      columns[1],
			"content":    columns[2],
			"created_at": columns[3],
		}
		noteList = append(noteList, note)
	}

	file, err := os.Create("notes.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(noteList)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Notes exported to notes.json")
}

func exportToTXT(notes string) {
	file, err := os.Create("notes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(notes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Notes exported to notes.txt")
}
