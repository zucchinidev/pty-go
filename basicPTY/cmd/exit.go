package cmd

import (
	"io"
	"pty-go/basicPTY/errors"
	"pty-go/basicPTY/printer"
)

func exit() Cmd {
	return Cmd{
		Name: "exit",
		Help: "Program ends",
		Action: func(w io.Writer, args ...string) error {
			printer.Print(w, "Goodbye!! :)")
			return errors.NewExitError(0)
		},
	}
}
