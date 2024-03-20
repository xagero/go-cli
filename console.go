package cli

import (
	"context"
	"fmt"
	"github.com/xagero/go-cli/command"
)

type Console struct {
	name           string
	description    string
	version        string
	defaultCommand *command.Command
	commands       map[string]*command.Command
}

// Construct return new Console
func Construct(name string, description string, version string) *Console {

	console := new(Console)
	console.name = name
	console.description = description
	console.version = version
	console.commands = make(map[string]*command.Command)

	return console
}

// SetDefaultCommand set command by default, return Console
func (c *Console) SetDefaultCommand(defaultCommand *command.Command) *Console {
	c.defaultCommand = defaultCommand
	return c
}

// AddCommand add command.Command to Console
func (c *Console) AddCommand(cmd *command.Command) *Console {
	name := cmd.GetName()
	c.commands[name] = cmd
	return c
}

// Run execute console command
func (c *Console) Run(context context.Context, args []string) error {

	fmt.Println(args)

	if len(args) < 2 {
		if c.defaultCommand != nil {
			return c.defaultCommand.Run(context, args)
		}
		return nil
	} else {
		name := args[1]
		if cmd, ok := c.commands[name]; ok {
			return cmd.Run(context, args)
		} else {
			// @todo command "name" not found
			panic("command not found")
		}
	}

	return nil
}
