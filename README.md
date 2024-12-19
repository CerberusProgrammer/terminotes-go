# Terminotes

Terminotes is a simple note-taking application that runs in the terminal.

## Why Terminotes?

I'm full against the idea of using a note-taking application that requires me to create an account or sign in. I don't want my notes to be stored on a server somewhere. I want them to be stored locally on my machine. That's why I created Terminotes. It's a simple note-taking application that runs in the terminal and stores notes locally on your machine. No need to create an account or sign in.

Also I code on my tablet or phone, and my keyboard is connected, and I need to add a super fast note without opening a browser or an app that requires a lot of resources (And time to open).

So Terminotes is a simple solution for me, and I hope it will be for you too. A very simple terminal app but with the idea of being fast and easy to use.

## MVP

- Create notes with a title and content (body) in the terminal
- List all notes
- View a note by its title
- Delete a note by its title
- Search for notes by title or content
- Export notes to a JSON file
- Import notes from a JSON file
- Configuration file to store the notes file path
- An option to clear all notes
- An option to reset the configuration
- An option to show the configuration
- An option to show the configuration file path
- A help message
- A version message
- A configuration file
- Notes storage in SQLite
- Multiline notes support (Title and content)
- [Required] A CLI
- [Optional] Animation when adding, deleting, or viewing a note (I'm not sure about this one)
- [Optional] A GUI (I'm not sure about this one too)
- [Future] A book to store notes in different categories
- [Future] A reminder system
- [Future] Tags for notes
- [Future] Connection to a cloud service (Like Firebase) to store notes
- [Future] Sync notes between devices
- [Future] Encryption for notes
- [Future] Password protection for notes
- [Future] Markdown support

## MVP Usage

- `tn add` to add a note with only content (The title will be the first line of the content)
- `tn add -t "Title"` to add a note with a title and no content
- `tn add -c "Content"` to add a note with no title and only content
- `tn add -t "Title" -c "Content"` to add a note with a title and content
- `tn add -m` to add a note with a multiline content (The title will be the first line of the content) (The content will be added line by line until you type `[end]` in a new line)
- `tn add -t "Title" -m` to add a note with a title and multiline content (The content will be added line by line until you type `[end]` in a new line)
- `tn add -c "Content" -m` to add a note with no title and multiline content (The content will be added line by line until you type `[end]` in a new line)
- `tn add -t "Title" -c "Content" -m` to add a note with a title, content, and multiline content (The content will be added line by line until you type `[end]` in a new line)
- `tn list` to list all notes
- `tn view -id {id}` to view a note by its title
- `tn delete -id {id}` to delete a note by its title (A cursor selection will be shown to select the note if there are multiple notes with the same title)
- `tn search -q "Query"` to search for notes by title or content
- `tn help` to show the help message
- `tn version` to show the version of Terminotes
- `tn export` to export all notes to a JSON file
- `tn import` to import notes from a JSON file
- `tn config` to configure Terminotes
- `tn config -r` to reset the configuration
- `tn config -l` to show the configuration
- `tn config -s` to show the configuration file path
