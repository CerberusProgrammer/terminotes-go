# Terminotes

Terminotes is a simple note-taking application that runs in the terminal.

(Want to watch the cool website? [Go here](https://terminotes.netlify.app/)

## Why Terminotes?

I'm full against the idea of using a note-taking application that requires me to create an account or sign in. I don't want my notes to be stored on a server somewhere. I want them to be stored locally on my machine. That's why I created Terminotes. It's a simple note-taking application that runs in the terminal and stores notes locally on your machine. No need to create an account or sign in.

Also I code on my tablet or phone, and my keyboard is connected, and I need to add a super fast note without opening a browser or an app that requires a lot of resources (And time to open).

So Terminotes is a simple solution for me, and I hope it will be for you too. A very simple terminal app but with the idea of being fast and easy to use.

## Instalation

### Windows

Download the latest release from the [releases page](https://github.com/CerberusProgrammer/terminotes-go/releases) and unzip it, then run the installer.

### Custom build

#### Custom build in Windows

```bash
go build -o terminotes.exe ./src/main.go
mv terminotes.exe ../package
cd package
Compress-Archive -Path terminotes.exe, install.bat, uninstall.bat -DestinationPath terminotes_installer.zip
```

Once you have the `terminotes_installer.zip` file, you can do whatever you want with it.

## MVP

- [x] Create notes with a title and content (body) in the terminal
- [x] List all notes
- [x] View a note by its title
- [x] Delete a note by its title
- [x] Search for notes by title or content
- [x] Export notes to a JSON file
- [x] Import notes from a JSON file
- [x] Configuration file to store the notes file path
- [x] An option to clear all notes
- [x] A help message
- [x] Notes storage in SQLite # <-- Issue: Using os.execute.
- [x] Multiline notes support (Title and content)
- [x] [Required] A CLI
- [ ] [Required] An easy Windows installer
- [ ] [Optional] Animation when adding, deleting, or viewing a note (I'm not sure about this one)
- [ ] [Optional] A GUI (I'm not sure about this one too)
- [ ] [Future] A book to store notes in different categories
- [ ] [Future] A reminder system
- [ ] [Future] Tags for notes
- [ ] [Future/Optional] Connection to a cloud service (Like Firebase) to store notes
- [ ] [Future] Sync notes between devices
- [ ] [Future] Encryption for notes
- [ ] [Future] Password protection for notes
- [ ] [Future] Markdown support

## MVP Usage

- [x] `tn add` to add a note with only content (The title will be the first line of the content)
- [x] `tn add -t "Title"` to add a note with a title and no content
- [x] `tn add -c "Content"` to add a note with no title and only content
- [x] `tn add -t "Title" -c "Content"` to add a note with a title and content
- [x] `tn add -m` to add a note with a multiline content (The title will be the first line of the content) (The content will be added line by line until you type `[end]` in a new line)
- [x] `tn add -t "Title" -m` to add a note with a title and multiline content (The content will be added line by line until you type `[end]` in a new line)
- [x] `tn add -c "Content" -m` to add a note with no title and multiline content (The content will be added line by line until you type `[end]` in a new line)
- [x] `tn add -t "Title" -c "Content" -m` to add a note with a title, content, and multiline content (The content will be added line by line until you type `[end]` in a new line)
- [x] `tn list` to list all notes
- [x] `tn view -id {id}` to view a note by its title
- [x] `tn delete -id {id}` to delete a note by its title (A cursor selection will be shown to select the note if there are multiple notes with the same title)
- [x] `tn search -q "Query"` to search for notes by title or content
- [x] `tn help` to show the help message
- [x] `tn export` to export all notes to a JSON file
- [x] `tn import` to import notes from a JSON file
