package commands

import (
	"fmt"
	"io"
	"os"
	"pty-go/basicPTY/printer"
	"pty-go/basicPTY/scanner"
)

func init() {
	_ = Register(Base{
		Name:   "sf",
		Help:   "Shows the content of a file, we need the path",
		Action: showFile,
	})
}

func showFile(_ io.Reader, output io.Writer, scanner scanner.ArgsScanner) (exit bool) {
	if scanner.Len() != 1 {
		printer.Print(output, "Please specify one file!")
		return
	}

	args := scanner.Args()
	f, err := os.Open(args[0])
	if err != nil {
		printer.Print(output, "Cannot open %s: %s\n", args[0], err)
	}
	defer f.Close()

	if _, err := io.Copy(output, f); err != nil {
		printer.Print(output, "Cannot print %s: %s\n", args[0], err)
		return
	}
	_, _ = fmt.Fprintln(output)
	return
}
