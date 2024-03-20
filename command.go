package cli

import (
	"fmt"
	"os"
)

type Callback func() error

// Command simple command structure
type Command struct {
	name        string
	description string
	callback    Callback
}

// Construct return new Command
func Construct(name, description string) *Command {
	cmd := new(Command)
	cmd.name = name
	cmd.description = description

	return cmd
}

// SetCallback set callback function
func (c *Command) SetCallback(callback Callback) {
	c.callback = callback
}

// Run execute callback function
func (c *Command) Run(args ...string) error {
	if len(args) == 0 {
		args = os.Args[1:]
	}

	if c.callback != nil {
		return c.callback()
	}

	return nil
}

func (c *Command) PrintHelp() {
	fmt.Println("I am help")
}
