package command

import (
	"context"
	"fmt"
	"strings"

	"github.com/xagero/go-helper/helper"
)

type Callback func() error

// Command simple command structure
type Command struct {
	// config
	name        string
	description string
	config      map[string]string

	// callback fn
	callback       Callback
	callbackBefore Callback
	callbackAfter  Callback

	// cli parameters
	arguments map[string]*CmdArgument
	options   map[string]*CmdOption
}

// Construct return new Command
func Construct(name, desc string) *Command {

	// Create new command
	cmd := new(Command)
	cmd.name = name
	cmd.description = desc

	// Configure
	cmd.config = make(map[string]string)
	cmd.config["common_options"] = "Y"

	// Make arguments and options
	cmd.arguments = make(map[string]*CmdArgument)
	cmd.options = make(map[string]*CmdOption)

	return cmd
}

// NewCommand (alias for Construct) return new Command
func NewCommand(name, description string) *Command {
	return Construct(name, description)
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

// SetCallbackBefore set before callback function
func (cmd *Command) SetCallbackBefore(callback Callback) *Command {
	cmd.callbackBefore = callback
	return cmd
}

// SetCallbackAfter set after callback function
func (cmd *Command) SetCallbackAfter(callback Callback) *Command {
	cmd.callbackAfter = callback
	return cmd
}

// RunBefore execute before callback function
func (cmd *Command) RunBefore(context context.Context) error {
	if cmd.callbackBefore != nil {
		return cmd.callbackBefore()
	}

	return nil
}

// Run execute callback function
func (cmd *Command) Run(context context.Context, args []string) error {
	if cmd.callback != nil {
		return cmd.callback()
	}

	return nil
}

// RunAfter execute after callback function
func (cmd *Command) RunAfter(context context.Context) error {
	if cmd.callbackAfter != nil {
		return cmd.callbackAfter()
	}

	return nil
}

// EnableCommonOptions enable common options
func (cmd *Command) EnableCommonOptions() {
	cmd.config["common_options"] = "Y"
}

// DisableCommonOptions disable common options
func (cmd *Command) DisableCommonOptions() {
	cmd.config["common_options"] = "N"
}

// PrintHelp print command help
func (cmd *Command) PrintHelp() {

	fmt.Println("Description:")
	fmt.Println(strings.Repeat(" ", 2) + cmd.description)

	fmt.Println("\nUsage:")
	usage := cmd.name
	if len(cmd.options) > 0 {
		usage += " [options]"
	}
	fmt.Println(strings.Repeat(" ", 2) + usage)

	if len(cmd.arguments) > 0 {
		cmd.printHelpArguments()
	}
	if len(cmd.options) > 0 {
		cmd.printHelpOptions()
	}

}

func (cmd *Command) printHelpOptions() {
	fmt.Println("\nOptions:")

	var n, s, d string

	for _, option := range cmd.options {

		n = option.name
		d = option.description
		s = option.short

		switch option.Input() {
		case OptionValueNone:
			if helper.IsBlank(s) {
				fmt.Printf("\t--%s \t- %s\n", n, d)
			} else {
				fmt.Printf("\t--%s, -%s \t- %s\n", n, s, d)
			}
		case OptionValueOptional:
			fmt.Printf("\t--%s \t- %s\n", n, d)
		case OptionValueRequire:
			fmt.Printf("\t--%s \t- %s\n", n, d)
		}
	}
}

func (cmd *Command) printHelpArguments() {
	fmt.Println("\nArguments:")
	for _, argument := range cmd.arguments {

		var required string
		if argument.input == ArgumentRequired {
			required = "(*required)"
		} else if argument.input == ArgumentOptional {
			required = "(optional)"
		}
		fmt.Printf("\t%s - %s %s\n", argument.name, argument.description, required)
	}
}

// AddOption add command option
func (cmd *Command) AddOption(name, input, description string) *CmdOption {
	opt := new(CmdOption)
	opt.name = name
	opt.input = input
	opt.description = description

	opt.exists = false
	opt.value = ""

	cmd.options[name] = opt

	return opt
}
