package commands

import (
	"io"
	"pty-go/basicPTY/printer"
	"pty-go/basicPTY/scanner"
)

func init() {
	_ = Register(Base{
		Name: "exit",
		Help: "Program ends",
		Action: func(input io.Reader, output io.Writer, scanner scanner.ArgsScanner) (exit bool) {
			printer.Print(output, "Goodbye!! :)")
			return true
		},
	})
}
