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
		_, _ = fmt.Fprint(w, "You wrote \"")
		_, _ = w.Write(s.Bytes())
		_, _ = fmt.Fprint(w, "\"\n")
	}
}
