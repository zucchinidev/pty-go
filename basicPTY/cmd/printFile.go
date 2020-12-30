package cmd

import (
	"fmt"
	"io"
	"os"
	"pty-go/basicPTY/printer"
)

func shuffle() Cmd {
	return Cmd{
		Name: "sf",
		Help: "Shows the content of a file, we need the path",
		Action: func(w io.Writer, args ...string) error {
			if len(args) != 1 {
				printer.Print(w, "Please specify one file!")
				return nil
			}

			f, err := os.Open(args[0])
			if err != nil {
				printer.Print(w, "Cannot open %s: %s\n", args[0], err)
			}
			defer f.Close()

			if _, err := io.Copy(w, f); err != nil {
				printer.Print(w, "Cannot print %s: %s\n", args[0], err)
				return err
			}
			_, _ = fmt.Fprintln(w)
			return nil
		},
	}
}
