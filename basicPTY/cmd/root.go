package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"pty-go/basicPTY/commands"
	coreErrors "pty-go/basicPTY/errors"
	"pty-go/basicPTY/printer"
	"pty-go/basicPTY/scanner"
)

func Run() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	sayHello(w)
	a := scanner.NewArgsScanner()
	b := bytes.NewBuffer(nil)
	availableCommands := commands.Commands()
	for {
		a.Reset()
		b.Reset()
		for {
			s.Scan()
			b.Write(s.Bytes())
			extra := a.Parse(b)

			if extra == "" {
				break
			}
			b.WriteString(extra)
		}

		desiredCmd := a.CMD()
		if desiredCmd == "" {
			desiredCmd = "help"
		}

		found := false
		for _, availableCmd := range availableCommands {
			if availableCmd.Match(desiredCmd) {
				found = true
				err := availableCmd.Run(w, a.Args()...)
				checkError(err)
				break
			}
		}
		if !found {
			printer.Print(w, "%q not found. Use `help` for available commands\n", desiredCmd)
		}
	}
}

func sayHello(w io.Writer) {
	msg := `
		Some welcome message!!
		Introduce your command:
`
	printer.Print(w, msg)
}

func checkError(err error) {
	switch {
	default:
		var exit *coreErrors.Exit
		if errors.As(err, &exit) {
			os.Exit(exit.Code)
		}
	}
}
