package command

import (
	"context"
	"fmt"
	"github.com/xagero/go-helper/helper"
	"regexp"
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

// EnableCommonOptions enable common options
func (cmd *Command) EnableCommonOptions() {
	// @todo code me
}

// DisableCommonOptions disable common options
func (cmd *Command) DisableCommonOptions() {
	// @todo code me
}

// PrintHelp print command help
func (cmd *Command) PrintHelp() {

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

// GetOption return CmdOption
func (cmd *Command) GetOption(key string) *CmdOption {
	if opt, ok := cmd.options[key]; ok {
		return opt
	}
	return nil
}

// Name return CmdOption name
func (opt CmdOption) Name() string {
	return opt.name
}

func (opt CmdOption) Input() string {
	return opt.input
}

func (opt CmdOption) Description() string {
	return opt.description
}

// Exists return CmdOption exists
func (opt CmdOption) Exists() bool {
	return opt.exists
}

// Value return CmdOption value
func (opt CmdOption) Value() string {
	return opt.value
}

func (opt CmdOption) Short() string {
	return opt.short
}

func (cmd *Command) ListOptions() map[string]*CmdOption {
	return cmd.options
}

// SetOptionExists set if option exists in console
func (cmd *Command) SetOptionExists(key string, b bool) {
	if opt, ok := cmd.options[key]; ok {
		opt.exists = b
	} else {
		// @todo fallback OptionNotExistsFallback
		panic("Option " + key + " not exist")
	}
}

// SetOptionValue set option value if option exists in console
func (cmd *Command) SetOptionValue(key string, value string) {
	if opt, ok := cmd.options[key]; ok {
		if opt.exists {
			opt.value = value
		}
	} else {
		// @todo fallback OptionNotExistsFallback
		panic("Option " + key + " not exist")
	}
}

func (opt *CmdOption) SetShortSyntax(short string) {
	if opt.input != OptionValueNone {
		panic("Short syntax is for option_value_none")
	}

	if helper.IsBlank(short) || len(short) > 1 {
		panic("Invalid command option short syntax")
	}

	bytes := []byte(short)
	if match, _ := regexp.Match(`[a-z]`, bytes); !match {
		panic("Invalid command option short syntax")
	}

	opt.short = short
}
