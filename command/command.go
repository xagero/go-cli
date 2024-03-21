package command

import (
	"context"
	"fmt"
)

// Construct return new Command
func Construct(name, description string) *Command {
	cmd := new(Command)
	cmd.name = name
	cmd.description = description
	cmd.arguments = make(map[string]*CmdArgument)
	cmd.options = make(map[string]*CmdOption)

	return cmd
}

// GetName return Command name
func (c *Command) GetName() string {
	return c.name
}

// SetCallback set callback function
func (c *Command) SetCallback(callback Callback) *Command {
	c.callback = callback
	return c
}

func (c *Command) SetCallbackBefore(before Callback) *Command {
	c.callbackBefore = before
	return c
}

func (c *Command) SetCallbackAfter(after Callback) *Command {
	c.callbackAfter = after
	return c
}

// Run execute callback function
func (c *Command) Run(context context.Context, args []string) error {
	if c.callback != nil {
		return c.callback()
	}

	return nil
}

func (c *Command) PrintHelp() {
	fmt.Println("I am help")
}

// AddArgument add command argument
func (c *Command) AddArgument(name, inputArgument, description string) *Command {

	arg := new(CmdArgument)
	arg.name = name
	arg.input = inputArgument
	arg.description = description

	c.arguments[name] = arg

	return c
}

// AddOption add command option
func (c *Command) AddOption(name, inputArgument, description string) *Command {
	opt := new(CmdOption)
	opt.name = name
	opt.input = inputArgument
	opt.description = description

	c.options[name] = opt

	return c
}
