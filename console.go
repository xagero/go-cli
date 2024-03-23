package cli

import (
	"context"
	"log/slog"
	"strings"

	"github.com/xagero/go-cli/command"
)

type Console struct {

	// Config
	name        string
	description string
	version     string

	// Commands
	defaultCommand *command.Command
	commands       map[string]*command.Command
}

// Construct return new Console
func Construct(name, desc, version string) *Console {

	// Create new console
	console := new(Console)
	console.name = name
	console.description = desc
	console.version = version

	// Make commands
	console.commands = make(map[string]*command.Command)

	return console
}

// NewConsole (alias for Construct) return new Console
func NewConsole(name, desc, version string) *Console {
	return Construct(name, desc, version)
}

// SetDefaultCommand set command by default, return Console
func (console *Console) SetDefaultCommand(cmd *command.Command) *Console {
	console.defaultCommand = cmd
	return console
}

// AddCommand add command.Command to Console
func (console *Console) AddCommand(cmd *command.Command) *Console {
	name := cmd.GetName()
	console.commands[name] = cmd
	return console
}

// Run execute console command
func (console *Console) Run(context context.Context, args []string) error {

	console.applyBuiltinFeatures()

	if len(args) < 2 {
		if console.defaultCommand == nil {
			console.PrintBanner()
			console.PrintHelp()
		} else {
			slog.Debug("Default command call")
			return console.defaultCommand.Run(context, args)
		}
	} else {
		name := args[1]
		if cmd, ok := console.commands[name]; ok {

			a := args[2:]
			console.processArguments(a, cmd)
			console.processOptions(a, cmd)

			cmd.ValidateArgumentRequirement()
			cmd.ValidateOptionRequirement()

			return cmd.Run(context, args)
		} else {
			console.PrintBanner()
			console.PrintHelp()
		}
	}

	return nil
}

func (console *Console) processArguments(a []string, cmd *command.Command) {
	idx := -1
	for _, value := range a {
		if strings.HasPrefix(value, "-") {
			continue // skip option
		}

		idx++
		cmd.SetArgumentValue(idx, value)
	}
}

// processOptions
func (console *Console) processOptions(a []string, cmd *command.Command) {
	for _, value := range a {
		if strings.HasPrefix(value, "--") {
			if strings.Contains(value, "=") {
				parts := strings.Split(value, "=")

				k := strings.TrimPrefix(parts[0], "--")
				v := strings.TrimPrefix(parts[1], `"`)

				cmd.SetOptionExist(k, true)
				cmd.SetOptionValue(k, v)

			} else {
				k := strings.TrimPrefix(value, "--")
				cmd.SetOptionExist(k, true)
			}
		} else if strings.HasPrefix(value, "-") {
			short := strings.TrimPrefix(value, "-")
			cmd.SetOptionExistByShort(short, true)
		}
	}
}

func (console *Console) applyBuiltinFeatures() {
	// @todo builtin features - help, quiet and verbose
}
