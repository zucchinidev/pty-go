package main

import (
	"bufio"
	"errors"
	"os"
	"pty-go/basicPTY/cmd"
	coreErrors "pty-go/basicPTY/errors"
	"pty-go/basicPTY/printer"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	printer.Print(w, "Some welcome message\n")
	availableCmds := cmd.Commands()
	for {
		s.Scan()
		args := strings.Split(string(s.Bytes()), " ")
		desiredCmd := args[0]
		restArgs := []string{}
		if len(args) > 1 {
			restArgs = args[1:]
		}
		found := false
		for _, availableCmd := range availableCmds {
			if !availableCmd.Match(desiredCmd) {
				continue
			}
			found = true
			err := availableCmd.Run(w, restArgs...)
			checkError(err)
		}
		if !found {
			printer.Print(w, "%q not found. Use `help` for available commands\n", args[0])
		}
	}
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
