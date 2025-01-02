package importation

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	jsonmanager "terminotes/src/data/json-manager"
)

func ImportCommand(args []string) {
	importCmd := flag.NewFlagSet("import", flag.ExitOnError)
	format := importCmd.String("format", "", "Format to import (sql, json, txt)")
	filePath := importCmd.String("file", "", "Path to the file to import")

	importCmd.Parse(args)

	if *format == "" || *filePath == "" {
		fmt.Println("Format and file path are required")
		importCmd.PrintDefaults()
		return
	}

	data, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(*format) {
	case "sql":
		importFromSQL(string(data))
	case "json":
		importFromJSON(data)
	case "txt":
		importFromTXT(string(data))
	default:
		fmt.Println("Unsupported format. Supported formats are: sql, json, txt")
	}
}

func importFromSQL(data string) {
	queries := strings.Split(data, ";")
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		err := jsonmanager.ExecuteSQL(query)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Data imported from SQL file")
}

func importFromJSON(data []byte) {
	var notes []map[string]interface{}
	err := json.Unmarshal(data, &notes)
	if err != nil {
		log.Fatal(err)
	}

	for _, note := range notes {
		title, _ := note["title"].(string)
		content, _ := note["content"].(string)
		err := jsonmanager.CreateNoteWithoutID(title, content)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Data imported from JSON file")
}

func importFromTXT(data string) {
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		if len(parts) != 4 {
			continue
		}
		title := parts[1]
		content := parts[2]
		err := jsonmanager.CreateNoteWithoutID(title, content)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Data imported from TXT file")
}
