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
	argsScanner := scanner.NewArgsScanner()
	b := bytes.NewBuffer(nil)
	for {
		argsScanner.Reset()
		b.Reset()
		for {
			s.Scan()
			b.Write(s.Bytes())
			extra := argsScanner.Parse(b)

			if extra == "" {
				break
			}
			b.WriteString(extra)
		}

		desiredCmd := argsScanner.CMD()
		if desiredCmd == "" {
			desiredCmd = "help"
		}

		errExecuting := commands.GetCommand(desiredCmd).Run(nil, w, argsScanner)
		if errExecuting != nil {
			var exit *coreErrors.Exit
			if errors.As(errExecuting, &exit) {
				os.Exit(exit.Code)
			}
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
