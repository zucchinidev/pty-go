package main

import (
	"bufio"
	"bytes"
	"os"
	"pty-go/basicPTY/color"
	"pty-go/basicPTY/commands"
	"pty-go/basicPTY/scanner"
)

func Run() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	color.Magenta.Colour(w, "** Welcome to PseudoTerm!! **\nPlease enter a command:\n")
	commands.Startup(w)
	defer commands.Shutdown(w)
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

		if commands.GetCommand(desiredCmd).Run(nil, w, argsScanner) {
			return
		}
	}
}
