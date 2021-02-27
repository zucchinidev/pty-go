package commands

import (
	"fmt"
	"io"
	"math/rand"
	"pty-go/basicPTY/color"
	"pty-go/basicPTY/printer"
	"pty-go/basicPTY/scanner"
)

func init() {
	_ = Register(Base{
		Name:   "pf",
		Help:   "Shuffle the arguments given",
		Action: shuffleParams,
	})
}

func shuffleParams(_ io.Reader, output io.Writer, scanner scanner.ArgsScanner) (err error) {
	args := scanner.Args()
	rand.Shuffle(len(args), func(i, j int) { args[i], args[j] = args[j], args[i] })
	for i := range args {
		if i > 0 {
			printer.Print(output, " ")
		}

		var f func(w io.Writer, format string, args ...interface{})
		f = color.Green.Colour
		if i%2 == 0 {
			f = color.Red.Colour
		}

		f(output, "%s", args[i])
	}
	_, _ = fmt.Fprintln(output)
	return nil
}
