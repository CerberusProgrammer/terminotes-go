package add

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	sqlitemanager "terminotes/src/data/sqlite-manager"
)

func AddCommand(args []string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	title := addCmd.String("t", "Untitled", "Title of the note (optional)")
	content := addCmd.String("c", "", "Content of the note (required)")
	multiline := addCmd.Bool("m", false, "Enable multiline input for the content")

	addCmd.Parse(args)

	if *multiline {
		fmt.Println("Enter your note content. Type ':wq' on a new line to save and exit.")
		*content = readMultilineInput()
	}

	if *content == "" {
		addCmd.PrintDefaults()
		return
	}

	id, err := sqlitemanager.CreateNote(*title, *content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Note %d created successfully\n", id)
}

func readMultilineInput() string {
	reader := bufio.NewReader(os.Stdin)
	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		if line == ":wq" {
			break
		}
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
