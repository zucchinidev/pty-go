package cmd

import (
	"fmt"
	"io"
	"math/rand"
	"pty-go/basicPTY/printer"
)

func printFile() Cmd {
	return Cmd{
		Name: "pf",
		Help: "Shuffle the arguments given",
		Action: func(w io.Writer, args ...string) error {
			rand.Shuffle(len(args), func(i, j int) { args[i], args[j] = args[j], args[i] })
			for i := range args {
				if i > 0 {
					printer.Print(w, " ")
				}
				printer.Print(w, "%s", args[i])
			}
			_, _ = fmt.Fprintln(w)
			return nil
		},
	}
}
