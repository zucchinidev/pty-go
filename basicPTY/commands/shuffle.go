package commands

import (
	"fmt"
	"io"
	"math/rand"
	"pty-go/basicPTY/color"
	"pty-go/basicPTY/printer"
)

func shuffle() Cmd {
	return Cmd{
		Name: "pf",
		Help: "Shuffle the arguments given",
		Action: func(w io.Writer, args ...string) error {
			rand.Shuffle(len(args), func(i, j int) { args[i], args[j] = args[j], args[i] })
			for i := range args {
				if i > 0 {
					printer.Print(w, " ")
				}

				var f func(w io.Writer, format string, args ...interface{})
				f = color.Green.Colour
				if i%2 == 0 {
					f = color.Red.Colour
				}

				f(w, "%s", args[i])
			}
			_, _ = fmt.Fprintln(w)
			return nil
		},
	}
}
