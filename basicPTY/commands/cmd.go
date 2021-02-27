package commands

import (
	"errors"
	"github.com/agnivade/levenshtein"
	"io"
	"pty-go/basicPTY/printer"
	"pty-go/basicPTY/scanner"
	"strings"
)

type Command interface {
	GetName() string
	GetHelp() string
	Run(input io.Reader, output io.Writer, scanner scanner.ArgsScanner) (err error)
}

// ErrDuplicateCommand is returned when two commands have the same name
var ErrDuplicateCommand = errors.New("Duplicate command")

var commands []Command

func Register(command Command) error {
	for i, c := range commands {
		// unique commands in alphabetical order
		switch strings.Compare(c.GetName(), command.GetName()) {
		case 0:
			return ErrDuplicateCommand
		case 1:
			commands = append(commands, nil)
			copy(commands[i+1:], commands[i:])
			commands[i] = command
			return nil
		case -1:
			continue
		}
	}
	commands = append(commands, command)
	return nil
}

// Base is a basic Command that runs a closure
type Base struct {
	Name, Help string
	Action     func(input io.Reader, output io.Writer, scanner scanner.ArgsScanner) (err error)
}

func (b Base) String() string { return b.Name }

// GetName returns the Name
func (b Base) GetName() string { return b.Name }

// GetHelp returns the Help
func (b Base) GetHelp() string { return b.Help }

// Run calls the closure
func (b Base) Run(input io.Reader, output io.Writer, scanner scanner.ArgsScanner) (err error) {
	return b.Action(input, output, scanner)
}

var suggest = Base{
	Action: func(input io.Reader, output io.Writer, scanner scanner.ArgsScanner) (err error) {
		var list []string
		for _, availableCmd := range commands {
			distance := levenshtein.ComputeDistance(availableCmd.GetName(), scanner.CMD())
			if distance < 3 {
				list = append(list, availableCmd.GetName())
			}
		}

		printer.Print(output, "%q not found. Use `help` for available commands\n", scanner.CMD())

		if len(list) > 0 {
			printer.Print(output, "Maybe you meant: %s\n", strings.Join(list, ", "))
		}

		return nil
	},
}

// GetCommand returns the command with the given name
func GetCommand(name string) Command {
	for _, c := range commands {
		if c.GetName() == name {
			return c
		}
	}
	return suggest
}
