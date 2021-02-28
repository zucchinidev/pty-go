package commands

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"pty-go/basicPTY/color"
	"pty-go/basicPTY/scanner"
)

func init() {
	_ = Register(&Stack{})
}

// The strings store in the Stack represent the state of the command
type Stack struct {
	data []string
}

func (s *Stack) GetName() string {
	return "stack"
}

func (s *Stack) GetHelp() string {
	return "a stack-like memory storage"
}

// There will be two sub-commands:
// push: followed by one or more arguments, will push to the stack.
// pop: will take the topmost element of the stack and it will not need any argument.
func (s *Stack) Run(_ io.Reader, output io.Writer, scanner scanner.ArgsScanner) (exit bool) {
	if scanner.LenArgs() < 1 || !s.isValid(scanner) {
		color.Red.Colour(output, "Use `stack push <something>` or `stack pop`\n")
		return
	}
	args := scanner.Args()

	subcommand := args[0]
	if subcommand == "push" {
		s.push(args[1:]...)
	}

	if subcommand == "pop" {
		v, ok := s.pop()
		if !ok {
			color.Cyan.Colour(output, "Stack Empty!!\n")
			return
		}
		color.Cyan.Colour(output, "Got: `%s`\n", v)
	}
	return
}

func (s *Stack) getPath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(u.HomeDir, ".stack"), nil
}

func (s *Stack) Startup() error {
	path, err := s.getPath()
	if err != nil {
		return err
	}
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()
	s.data = s.data[:0]
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		s.push(string(sc.Bytes()))
	}
	return nil
}

func (s *Stack) Shutdown() error {
	path, err := s.getPath()
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, v := range s.data {
		if _, err := fmt.Fprintln(f, v); err != nil {
			return err
		}
	}
	return nil
}

func (s *Stack) isValid(scanner scanner.ArgsScanner) bool {
	args := scanner.Args()
	switch args[0] {
	case "pop":
		return len(args) == 1
	case "push":
		return len(args) > 1
	default:
		return false
	}
}

func (s *Stack) push(values ...string) {
	s.data = append(s.data, values...)
}

func (s *Stack) pop() (string, bool) {
	if len(s.data) == 0 {
		return "", false
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}
