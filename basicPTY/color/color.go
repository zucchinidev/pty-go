package color

import (
	"io"
	"pty-go/basicPTY/printer"
)

type color int

const (
	Reset   color = 0
	Red     color = 31
	Green   color = 32
	Yellow  color = 33
	Blue    color = 34
	Magenta color = 35
	Cyan    color = 36
	White   color = 37
)

func (c color) Start(w io.Writer) {
	printer.Print(w, "\x1b[%dm", c)
}

func (c color) End(w io.Writer) {
	printer.Print(w, "\x1b[%dm", Reset)
}

func (c color) Colour(w io.Writer, format string, args ...interface{}) {
	c.Start(w)
	printer.Print(w, format, args...)
	c.End(w)
}
