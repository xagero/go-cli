package command

import (
	"context"
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

// GetName (getter) return Command name
func (cmd *Command) GetName() string {
	return cmd.name
}

// GetDescription (getter) return Command description
func (cmd *Command) GetDescription() string {
	return cmd.description
}

// SetCallback set callback function
func (cmd *Command) SetCallback(callback Callback) *Command {
	cmd.callback = callback
	return cmd
}

func (cmd *Command) SetCallbackBefore(before Callback) *Command {
	cmd.callbackBefore = before
	return cmd
}

func (cmd *Command) SetCallbackAfter(after Callback) *Command {
	cmd.callbackAfter = after
	return cmd
}

// Run execute callback function
func (cmd *Command) Run(context context.Context, args []string) error {
	if cmd.callback != nil {
		return cmd.callback()
	}

	return nil
}

// PrintHelp print command help
func (cmd *Command) PrintHelp() {
	return
}

// AddArgument add command argument
func (cmd *Command) AddArgument(name, inputArgument, description string) *Command {

	arg := new(CmdArgument)
	arg.name = name
	arg.value = "" // An empty string
	arg.position = len(cmd.arguments)
	arg.input = inputArgument
	arg.description = description

	cmd.arguments[name] = arg

	return cmd
}

// AddOption add command option
func (cmd *Command) AddOption(name, inputArgument, description string) *Command {
	opt := new(CmdOption)
	opt.name = name
	opt.input = inputArgument
	opt.description = description

	cmd.options[name] = opt

	return cmd
}

// GetArgumentValue return argument value
func (cmd *Command) GetArgumentValue(name string) string {
	return cmd.arguments[name].value
}

// SetArgumentValue set argument value
func (cmd *Command) SetArgumentValue(position int, value string) *Command {
	for _, item := range cmd.arguments {
		if position == item.position {
			item.value = value
		}
	}

	return cmd
}
