package command

import "github.com/xagero/go-helper/helper"

// CmdArgument command argument, always ordered in console
type CmdArgument struct {
	// config
	name        string
	position    int
	input       string
	description string

	// cli input
	value string
}

// AddArgument add command argument
func (cmd *Command) AddArgument(name, input, description string) *CmdArgument {

	arg := new(CmdArgument)
	arg.name = name
	arg.value = "" // An empty string
	arg.position = len(cmd.arguments)
	arg.input = input
	arg.description = description

	cmd.arguments[name] = arg

	return arg
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

func (cmd *Command) ValidateArgumentRequirement() error {

	for _, argument := range cmd.arguments {
		if argument.input == ArgumentRequired {
			if helper.IsBlank(argument.value) {
				// @todo InvalidArgumentFallback
				panic("Invalid argument [ " + argument.name + " ], not exists")
			}
		}
	}

	return nil
}
