package cmd

import "io"

type FnCmd func(w io.Writer, args ...string) error

type Cmd struct {
	Name   string // the command name
	Help   string // the command description
	Action FnCmd
}

func (c Cmd) Match(s string) bool {
	return c.Name == s
}

func (c Cmd) Run(w io.Writer, args ...string) error {
	return c.Action(w, args...)
}

func Commands() []Cmd {
	return []Cmd{
		exit(),
		help(),
		printFile(),
		shuffle(),
	}
}
