package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	_, _ = fmt.Fprint(w, "Some welcome message\n")
	for {
		s.Scan()
		command := s.Bytes()
		if exit(command) {
			os.Exit(0)
		}
		_, _ = fmt.Fprintf(w, "You wrote %q\n", command)
	}
}

func exit(command []byte) bool {
	return string(command) == "exit"
}
