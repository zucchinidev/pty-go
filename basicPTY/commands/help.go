package commands

import (
	"io"
	"pty-go/basicPTY/printer"
)

func help() Cmd {
	return Cmd{
		Name: "help",
		Help: "Shows available commands",
		Action: func(w io.Writer, args ...string) error {
			printer.Print(w, "Available commands:\n")
			for _, c := range Commands() {
				printer.Print(w, " - %-15s %s\n", c.Name, c.Help)
			}
			return nil
		},
	}
}
