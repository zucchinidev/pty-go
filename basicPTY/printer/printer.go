package printer

import (
	"fmt"
	"io"
)

func Print(w io.Writer, format string, args ...interface{}) {
	_, _ = fmt.Fprintf(w, format, args...)
}
