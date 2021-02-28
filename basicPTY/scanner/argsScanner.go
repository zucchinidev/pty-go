package scanner

import (
	"bufio"
	"io"
)

type ArgsScanner interface {
	Len() int
	Args() []string
	LenArgs() int
	CMD() string
}

type argsScanner []string

func NewArgsScanner() *argsScanner {
	return &argsScanner{}
}

func (a *argsScanner) Reset() { *a = (*a)[0:0] }

func (a *argsScanner) Parse(r io.Reader) (extra string) {
	s := bufio.NewScanner(r)
	s.Split(ScanArgs)

	for s.Scan() {
		*a = append(*a, s.Text())
	}

	if len(*a) == 0 {
		return ""
	}

	lastArg := (*a)[len(*a)-1]

	if !isQuote(rune(lastArg[0])) {
		return ""
	}

	*a = (*a)[:len(*a)-1]

	return lastArg + "\n"
}

func (a *argsScanner) Len() int {
	return len(*a)
}

func (a *argsScanner) Args() []string {
	if a.Len() > 1 {
		return (*a)[1:]
	}

	return []string{}
}

func (a *argsScanner) LenArgs() int {
	return len(a.Args())
}

func (a *argsScanner) CMD() string {
	if a.Len() == 0 {
		return ""
	}

	return (*a)[0]
}
