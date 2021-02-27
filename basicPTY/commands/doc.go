package commands

// Embedded unnamed field (inherits method)
/*

type MyCmd struct {
	Base
	MyField string
}

// custom implementation
type MyImpl struct{}

func (MyImpl) GetName() string { return "myimpl" }
func (MyImpl) GetHelp() string { return "help string"}
func (MyImpl) Run(input io.Reader, output io.Writer, args ...string) bool {
	// do something
	return true
}

The difference between the MyCmd struct and the MyImpl struct is that one can be used as decorator for another command,
while the second is a different implementation so it can't interact with another command.

*/
