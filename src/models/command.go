package models

type Command struct {
	Name        string
	Description string
	Execute     func(args []string)
	Error       string
}
