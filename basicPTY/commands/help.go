package commands

import (
	"io"
	"pty-go/basicPTY/printer"
	"pty-go/basicPTY/scanner"
)

func init() {
	_ = Register(Base{
		Name: "help",
		Help: "Shows available commands",
		Action: func(input io.Reader, output io.Writer, scanner scanner.ArgsScanner) (exit bool) {
			printer.Print(output, "Available commands:\n")
			for _, c := range commands {
				printer.Print(output, " - %-15s %s\n", c.GetName(), c.GetHelp())
			}
			return
		},
	})
}
